package utils

// Inspired by https://github.com/helm/helm/blob/10f880a876d030aab6ee8861c1a6b5f02fcbadd3/pkg/helm/portforwarder/portforwarder.go
// This has modifications to enable use to specify the labels.
import (
	"fmt"
	"log"

	"k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

// NewPortForward creates a new and initialized tunnel to a pod's port.
func NewPortForward(namespace string, labels labels.Set, port int, client kubernetes.Interface, config *rest.Config) (*Tunnel, error) {
	podName, err := GetPodName(client.CoreV1(), labels, namespace)
	if err != nil {
		return nil, err
	}
	t := NewTunnel(client.CoreV1().RESTClient(), config, namespace, podName, port)
	return t, t.ForwardPort()
}

// GetPodName fetches the name of pod meeting x labels
func GetPodName(client corev1.PodsGetter, labels labels.Set, namespace string) (string, error) {
	selector := labels.AsSelector()
	pod, err := getFirstRunningPod(client, namespace, selector)
	if err != nil {
		return "", err
	}

	return pod.ObjectMeta.GetName(), nil
}

// getFirstRunningPod returns a pod that meets our selector.
func getFirstRunningPod(client corev1.PodsGetter, namespace string, selector labels.Selector) (*v1.Pod, error) {
	log.Printf("searching for a pod with: '%s' in namespace '%s'", selector.String(), namespace)

	options := metav1.ListOptions{LabelSelector: selector.String()}
	pods, err := client.Pods(namespace).List(options)
	if err != nil {
		return nil, err
	}

	if len(pods.Items) < 1 {
		return nil, fmt.Errorf("could not find a pod meeting that criteria")
	}

	for _, p := range pods.Items {
		// TODO: check if ready.
		return &p, nil
	}
	return nil, fmt.Errorf("could not find a ready pod meeting that criteria")
}
