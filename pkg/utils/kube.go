package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"

	// for authentication
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"

	"k8s.io/client-go/rest"
)

// GetKubeConfig gets the config for a specified kube context
func GetKubeConfig(context string) (*rest.Config, error) {
	home := os.Getenv("HOME")

	if home == "" {
		return nil, errors.New("Failed to find kubeconfig (HOME undef)")
	}

	kubeConfig := filepath.Join(home, ".kube/config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}

	log.Printf("using host: %s", config.Host)
	return config, nil
}

// GetKubeClient creates a Kubernetes config and client for a given kubeconfig context.
func GetKubeClient(context string) (*rest.Config, kubernetes.Interface, error) {
	config, err := GetKubeConfig(context)
	if err != nil {
		return nil, nil, err
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, fmt.Errorf("could not get Kubernetes client: %s", err)
	}
	return config, client, nil
}
