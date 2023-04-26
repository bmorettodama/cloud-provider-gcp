package highperf

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	coreinformers "k8s.io/client-go/informers/core/v1"
	fakeClient "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/tools/cache"
	fakeNetwork "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned/fake"
	networkinformer "k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1"
	alphanetworkinformer "k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1alpha1"
	"k8s.io/cloud-provider-gcp/pkg/controller/testutil"
	"k8s.io/cloud-provider-gcp/providers/gce"
	"k8s.io/component-base/metrics/prometheus/controllers"
	"testing"
	"time"

	fakeSriovClient "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/fake"
)

type testHighPerfController struct {
	stop          context.CancelFunc
	ctx           context.Context
	gceCloud      *gce.Cloud
	sriovClient   *fakeSriovClient.Clientset
	networkClient *fakeNetwork.Clientset
	nwInformer    cache.SharedIndexInformer
	gnpInformer   cache.SharedIndexInformer
	nodeInformer  coreinformers.NodeInformer
	clusterValues gce.TestClusterValues
	controller    *Controller
	metrics       *controllers.ControllerManagerMetrics
}

func setupHighPerfController(ctx context.Context) *testHighPerfController {
	fakeNetworking := fakeNetwork.NewSimpleClientset()
	fakeSriov := fakeSriovClient.NewSimpleClientset()

	fakeNodeHandler := &testutil.FakeNodeHandler{
		Existing: []*v1.Node{
			{ObjectMeta: metav1.ObjectMeta{Name: "node0"}},
		},
		Clientset: fakeClient.NewSimpleClientset(),
	}
	fakeClient := &fakeClient.Clientset{}
	fakeInformerFactory := informers.NewSharedInformerFactory(fakeClient, 0*time.Second)
	nodeInformer := fakeInformerFactory.Core().V1().Nodes()

	for _, node := range fakeNodeHandler.Existing {
		nodeInformer.Informer().GetStore().Add(node)
	}

	nwInformer := networkinformer.NewNetworkInformer(fakeNetworking, 0*time.Second, cache.Indexers{})
	gnpInformer := alphanetworkinformer.NewGKENetworkParamSetInformer(fakeNetworking, 0*time.Second, cache.Indexers{})
	testClusterVals := gce.DefaultTestClusterValues()
	fakeGce := gce.NewFakeGCECloud(testClusterVals)

	controller := NewHighPerfController(
		fakeGce,
		fakeNetworking,
		nwInformer,
		gnpInformer,
		nodeInformer,
		fakeSriov,
	)
	metrics := controllers.NewControllerManagerMetrics("test")

	return &testHighPerfController{
		gceCloud:      fakeGce,
		networkClient: fakeNetworking,
		nwInformer:    nwInformer,
		gnpInformer:   gnpInformer,
		nodeInformer:  nodeInformer,
		clusterValues: testClusterVals,
		controller:    controller,
		metrics:       metrics,
	}
}

func (testVals *testHighPerfController) runHighPerfController(ctx context.Context) {
	go testVals.nwInformer.Run(ctx.Done())
	go testVals.gnpInformer.Run(ctx.Done())
	go testVals.controller.Run(1, ctx.Done(), testVals.metrics)
}

func TestControllerRuns(t *testing.T) {
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	testVals := setupHighPerfController(ctx)
	testVals.runHighPerfController(ctx)
}
