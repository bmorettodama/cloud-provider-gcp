package highperf

import (
	"context"
	"fmt"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	coreinformers "k8s.io/client-go/informers/core/v1"

	networkclientset "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned"
	"k8s.io/cloud-provider-gcp/providers/gce"
	controllersmetrics "k8s.io/component-base/metrics/prometheus/controllers"

	networkv1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1"
	networkv1alpha1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1alpha1"
)

const (
	numOfRetries = 5
)

type Controller struct {
	gceCloud         *gce.Cloud
	networkClientset networkclientset.Interface

	nwInformer   cache.SharedIndexInformer
	gnpInformer  cache.SharedIndexInformer
	nodeInformer coreinformers.NodeInformer

	queue workqueue.RateLimitingInterface
}

func NewHighPerfController(
	gceCloud *gce.Cloud,
	networkClientset networkclientset.Interface,
	nwInformer cache.SharedIndexInformer,
	gnpInformer cache.SharedIndexInformer,
	nodeInformer coreinformers.NodeInformer) *Controller {

	hc := &Controller{
		gceCloud:         gceCloud,
		networkClientset: networkClientset,
		nwInformer:       nwInformer,
		gnpInformer:      gnpInformer,
		nodeInformer:     nodeInformer,
		queue:            workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "highperf"),
	}

	// TODO (bmorettodama): Add event handlers for Network and GNP object DeleteFunc

	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    hc.addNode,
		UpdateFunc: hc.updateNode,
	})

	return hc
}

func (c *Controller) addNode(obj interface{}) {
	key, err := cache.MetaNamespaceKeyFunc(obj)
	if err != nil {
		klog.Errorf("failed to get node key: %v", err)
		return
	}
	c.queue.Add(key)
}

func (c *Controller) updateNode(old interface{}, new interface{}) {
	oldNode := old.(*v1.Node)
	newNode := new.(*v1.Node)

	networkStatusOld := oldNode.Annotations[networkv1.NodeNetworkAnnotationKey]
	networkStatusNew := newNode.Annotations[networkv1.NodeNetworkAnnotationKey]
	northInterfaceOld := oldNode.Annotations[networkv1.NorthInterfacesAnnotationKey]
	northInterfaceNew := newNode.Annotations[networkv1.NorthInterfacesAnnotationKey]

	if (networkStatusOld == networkStatusNew) && (northInterfaceOld == northInterfaceNew) {
		return
	}
	key, err := cache.MetaNamespaceKeyFunc(newNode)
	if err != nil {
		klog.Errorf("failed to add node key to worker queue: %v", err)
		return
	}
	c.queue.Add(key)
}

func (c *Controller) Run(numWorkers int, stopCh <-chan struct{}, controllerManagerMetrics *controllersmetrics.ControllerManagerMetrics) {
	defer utilruntime.HandleCrash()

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()
	defer c.queue.ShutDown()

	klog.Info("Starting high perf controller")
	defer klog.Info("Shutting down high perf controller")
	controllerManagerMetrics.ControllerStarted("highperf")
	defer controllerManagerMetrics.ControllerStopped("highperf")

	if !cache.WaitForNamedCacheSync("node", stopCh, c.nodeInformer.Informer().HasSynced) {
		return
	}

	if !cache.WaitForNamedCacheSync("networks", stopCh, c.nwInformer.HasSynced) {
		return
	}

	if !cache.WaitForNamedCacheSync("gkenetworkparamset", stopCh, c.gnpInformer.HasSynced) {
		return
	}

	for i := 0; i < numWorkers; i++ {
		go wait.UntilWithContext(ctx, c.runWorker, time.Second)
	}

	<-stopCh
}

func (c *Controller) runWorker(ctx context.Context) {
	for c.processNextItem(ctx) {
	}
}

func (c *Controller) processNextItem(ctx context.Context) bool {
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	defer c.queue.Done(key)

	err := c.processNode(ctx, key.(string))
	c.handleErr(err, key)
	return true
}

func (c *Controller) handleErr(err error, key interface{}) {
	if err == nil {
		c.queue.Forget(key)
		return
	}

	// This controller retries 5 times if something goes wrong. After that, it stops trying.
	if c.queue.NumRequeues(key) < numOfRetries {
		klog.Warningf("Error while updating GKENetworkParamSet object, retrying %v: %v", key, err)

		// Re-enqueue the key rate limited. Based on the rate limiter on the
		// queue and the re-enqueue history, the key will be processed later again.
		c.queue.AddRateLimited(key)
		return
	}

	c.queue.Forget(key)
	// Report to an external entity that, even after several retries, we could not successfully process this key
	utilruntime.HandleError(err)
	klog.Errorf("dropping node  %q out of the queue: %v", key, err)
}

func (c *Controller) processNode(ctx context.Context, key string) error {
	node, err := c.nodeInformer.Lister().Get(key)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return fmt.Errorf("fetching object with key %s from store failed with %v", key, err)
	}

	//northInterface and network-status annotations need to be ready before we move on
	if node.Annotations[networkv1.NorthInterfacesAnnotationKey] == "" || node.Annotations[networkv1.NodeNetworkAnnotationKey] == "" {
		return nil
	}

	// check if northInterface and network-status have entries for same network
	northInterfaces, err := networkv1.ParseNorthInterfacesAnnotation(node.Annotations[networkv1.NorthInterfacesAnnotationKey])
	if err != nil {
		return fmt.Errorf("parsing northInterface annotation: %v", err)
	}
	nodeNetworks, err := networkv1.ParseNodeNetworkAnnotation(node.Annotations[networkv1.NodeNetworkAnnotationKey])
	if err != nil {
		fmt.Errorf("parsing network-status annotation: %v", err)
	}
	var readyNetworks networkv1.NorthInterfacesAnnotation

	// create a list of networks present in both the northInterface and network-status annotations
	statusMap := make(map[string]struct{})
	for _, nodeNetwork := range nodeNetworks {
		statusMap[nodeNetwork.Name] = struct{}{}
	}
	for _, northInterface := range northInterfaces {
		if _, ok := statusMap[northInterface.Network]; ok {
			readyNetworks = append(readyNetworks, northInterface)
		}
	}

	// return when there are no ready netwoks
	if len(readyNetworks) == 0 {
		return nil
	}

	for _, network := range readyNetworks {
		// Fetch network object and check if its Device type
		obj, exists, err := c.nwInformer.GetIndexer().GetByKey(network.Network)
		if err != nil {
			return fmt.Errorf("fetching %s network object: %v", network.Network, err)
		}
		if !exists {
			continue
		}
		nwobj := obj.(*networkv1.Network)

		if nwobj.Spec.Type != networkv1.DeviceNetworkType {
			continue
		}

		klog.Info("Network %s is of type Device", network.Network)

		// Fetch GNP object for Device type networks
		obj, exists, err = c.gnpInformer.GetIndexer().GetByKey(nwobj.Spec.ParametersRef.Name)
		if err != nil {
			return fmt.Errorf("fetching %s gnp object: %v", nwobj.Spec.ParametersRef.Name, err)
		}
		if !exists {
			continue
		}
		gnpobj := obj.(*networkv1alpha1.GKENetworkParamSet)

		err = c.updateConfigMap(node, nwobj, gnpobj)
		if err != nil {
			return err
		}

		err = c.updateNodeStateSpec(node, nwobj, gnpobj)
		if err != nil {
			return err
		}
	}

	return nil
}

// TODO (bmorettodama): implement function to update config map for device type networks
func (c *Controller) updateConfigMap(nodeobj *v1.Node, nwobj *networkv1.Network, gnpobj *networkv1alpha1.GKENetworkParamSet) error {
	return nil
}

// TODO (bmorettodama): implement function to udpate node state for device type networks
func (c *Controller) updateNodeStateSpec(nodeobj *v1.Node, nwobj *networkv1.Network, gnpobj *networkv1alpha1.GKENetworkParamSet) error {
	return nil
}
