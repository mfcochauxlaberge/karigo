package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := karigo.NewServer()
	server.Config = config

	// Source
	// TODO Stop ignoring this configuration and
	// use the proper driver for the source.
	src := &memory.Source{}

	_ = src.Reset(karigo.FirstSchema())

	// Add cluster control node
	ctlNode := karigo.NewNode(config)
	ctlNode.Name = "main_node"

	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
