package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/jaredallard/k8s-mc-helper/pkg/proto/mchelper"

	"github.com/urfave/cli"
)

func runCommand(c *cli.Context) error {
	command := c.Args().First() + " " + strings.Join(c.Args().Tail(), " ")

	fmt.Printf("Executing command: '%s'\n", command)

	client, err := GetClient()
	if err != nil {
		return err
	}

	stream, err := client.SendCommand(context.Background(), &mchelper.SendCommandRequest{
		Command: command,
	})
	if err != nil {
		return err
	}

	// TODO: In the future we can send back output here
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}
	}

	fmt.Println("done")

	return nil
}
