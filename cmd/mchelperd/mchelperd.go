package main

import (
	"log"
	"net"

	"github.com/jaredallard/k8s-mc-helper/pkg/minecraft"
	"github.com/jaredallard/k8s-mc-helper/pkg/proto/mchelper"
	"google.golang.org/grpc"
)

func main() {
	server, err := minecraft.FindServer()
	if err != nil {
		log.Fatalf("Failed to find minecraft server: %s", err.Error())
	}
	log.Printf("minecraft is running: pid=%d,cwd=%s\n", server.PID, server.Dir)

	err = server.SendCommand("/say hi")
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}

	l, err := net.Listen("tcp", "0.0.0.0:4423")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := newmchelperd(server)
	mchelper.RegisterOperationsServicesServer(s, *srv)

	log.Printf("Started. Listening on port 4423")
	s.Serve(l)
}
