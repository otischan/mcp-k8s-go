package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKubeConfig() clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, nil)
	return kubeConfig
}

func GetCurrentContext() (string, error) {
	kubeConfig := GetKubeConfig()
	config, err := kubeConfig.RawConfig()
	if err != nil {
		return "", err
	}
	return config.CurrentContext, nil
}

func GetKubeClientset() (*kubernetes.Clientset, error) {
	kubeConfig := GetKubeConfig()

	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		return nil, err
	}
	return clientset, nil
}
