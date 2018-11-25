package main

import (
	"log"

	"github.com/jaredallard/k8s-mc-helper/pkg/minecraft"
	"github.com/jaredallard/k8s-mc-helper/pkg/proto/mchelper"
)

type srv struct {
	Minecraft *minecraft.Server
}

func newmchelperd(m *minecraft.Server) *srv {
	return &srv{
		Minecraft: m,
	}
}

func (s srv) ListMods(e *mchelper.Empty, srv mchelper.OperationsServices_ListModsServer) error {
	return nil
}

func (s srv) SendCommand(r *mchelper.SendCommandRequest, srv mchelper.OperationsServices_SendCommandServer) error {
	log.Printf("executing command: '%s'", r.Command)
	err := s.Minecraft.SendCommand(r.Command)
	return err
}
