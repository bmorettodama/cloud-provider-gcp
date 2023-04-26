package main

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/tools/cache"
	cloudprovider "k8s.io/cloud-provider"
	networkclientset "k8s.io/cloud-provider-gcp/crd/client/network/clientset/versioned"
	networkinformer "k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1"
	alphanetworkinformer "k8s.io/cloud-provider-gcp/crd/client/network/informers/externalversions/network/v1alpha1"
	highperfcontroller "k8s.io/cloud-provider-gcp/pkg/controller/highperf"
	"k8s.io/cloud-provider-gcp/providers/gce"
	"k8s.io/cloud-provider/app"
	cloudcontrollerconfig "k8s.io/cloud-provider/app/config"
	genericcontrollermanager "k8s.io/controller-manager/app"
	sriovclientset "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned"
	"k8s.io/controller-manager/controller"
)

func startHighPerfControllerWrapper(initCtx app.ControllerInitContext, config *cloudcontrollerconfig.CompletedConfig, cloud cloudprovider.Interface) app.InitFunc {
	return func(ctx context.Context, controllerCtx genericcontrollermanager.ControllerContext) (controller.Interface, bool, error) {
		return startHighPerfController(config, controllerCtx, cloud)
	}
}

func startHighPerfController(ccmConfig *cloudcontrollerconfig.CompletedConfig, controllerCtx genericcontrollermanager.ControllerContext, cloud cloudprovider.Interface) (controller.Interface, bool, error) {
	gceCloud, ok := cloud.(*gce.Cloud)
	if !ok {
		return nil, false, fmt.Errorf("HighPerfController does not support %v provider", cloud.ProviderName())
	}

	kubeConfig := ccmConfig.Complete().Kubeconfig
	kubeConfig.ContentType = "application/json" //required to serialize HighPerf objects to json
	networkClient, err := networkclientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, false, err
	}
	sriovClient, err := sriovclientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, false, err
	}

	nwInformer := networkinformer.NewNetworkInformer(networkClient, 0*time.Second, cache.Indexers{})
	gnpInformer := alphanetworkinformer.NewGKENetworkParamSetInformer(networkClient, 0*time.Second, cache.Indexers{})

	highPerfController := highperfcontroller.NewHighPerfController(
		gceCloud,
		networkClient,
		nwInformer,
		gnpInformer,
		controllerCtx.InformerFactory.Core().V1().Nodes(),
		sriovClient,
	)

	go nwInformer.Run(controllerCtx.Stop)
	go gnpInformer.Run(controllerCtx.Stop)

	go highPerfController.Run(1, controllerCtx.Stop, controllerCtx.ControllerManagerMetrics)

	return nil, true, nil
}
