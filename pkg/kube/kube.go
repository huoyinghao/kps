package kube

import (
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubeClient(kubeConfig string) kubernetes.Interface {
	conf, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}
	cli, err := kubernetes.NewForConfig(conf)
	if err != nil {
		panic(err)
	}
	return cli
}

func NewDynamicClient(kubeConfig string) dynamic.Interface {
	Config := kubeConfig
	var clusterConfig *rest.Config
	var err error
	if kubeConfig != "" {
		clusterConfig, err = clientcmd.BuildConfigFromFlags("", Config)
	} else {
		clusterConfig, err = rest.InClusterConfig()
	}
	if err != nil {
		panic(err)
	}
	clusterClient, err := dynamic.NewForConfig(clusterConfig)
	if err != nil {
		panic(err)
	}
	return clusterClient
}
