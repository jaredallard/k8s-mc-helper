package main

import (
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/labels"

	"github.com/jaredallard/k8s-mc-helper/pkg/utils"

	"github.com/jaredallard/k8s-mc-helper/pkg/proto/mchelper"
	"google.golang.org/grpc"
)

// GetClient returns a client to mchelperd via k8s
func GetClient() (mchelper.OperationsServicesClient, error) {

	host := os.Getenv("MCHELPERD_HOST")
	if os.Getenv("MCHELPERD_HOST") == "" {
		c, i, err := utils.GetKubeClient("")
		if err != nil {
			return nil, err
		}

		t, err := utils.NewPortForward("default", labels.Set{
			"app": "minecraft-minecraft",
		}, 4423, i, c)
		if err != nil {
			return nil, err
		}

		host = fmt.Sprintf("127.0.0.1:%d", t.Local)
	}

	// TODO: k8s
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := mchelper.NewOperationsServicesClient(conn)

	return client, nil
}
