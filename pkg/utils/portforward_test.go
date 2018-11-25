package utils

// TODO: migrate tiller/client -> utils.kube tests

// This is portforward.go tests but because kube and portforward pretty much overlap,
// we pass off testing for kubeclient to here.
import (
	"flag"
	"os"
	"testing"

	"k8s.io/apimachinery/pkg/labels"
)

var (
	integration = flag.Bool("integration", false, "run integration tests")
)

func TestMain(m *testing.M) {
	flag.Parse()

	if *integration {
		// could do something here to initialize for this test.
	}

	result := m.Run()

	os.Exit(result)
}

// This also covers kube.go
func TestPortForwardAll(t *testing.T) {
	if !*integration {
		t.Skip("not running integration tests")
		return
	}

	config, client, err := GetKubeClient("")
	if err != nil {
		t.Errorf("failed to create kubernetes client: %s", err.Error())
		return
	}

	podLabels := labels.Set{"k8s-app": "kube-dns"}
	podNamespace := "kube-system"

	podname, err := GetPodName(client.CoreV1(), podLabels, podNamespace)
	if podname == "" || err != nil {
		t.Error("failed to find a kube-dns pod, which means GetPodName is likely not working.")
		return
	}

	tun, err := NewPortForward(podNamespace, podLabels, 1000, client, config)
	if err != nil || tun == nil {
		t.Error("failed to create port-forward to a kube-dns pod")
		return
	}
}
