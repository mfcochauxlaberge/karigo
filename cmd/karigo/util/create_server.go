package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := &karigo.Server{
		Port:  config.Port,
		Nodes: map[string]*karigo.Node{},
	}

	// Journal
	var jrnl karigo.Journal

	switch config.Journal["driver"] {
	case "memory":
		jrnl = &memory.Journal{}
	case "psql":
		jrnl = &psql.Journal{}
	default:
		jrnl = &memory.Journal{}
	}

	// Source
	// TODO Stop ignoring this configuration and
	// use the proper driver for the source.
	src := &memory.Source{}

	_ = src.Reset()

	// Add cluster control node
	ctlNode := karigo.NewNode(jrnl, src)
	ctlNode.Name = "main_node"

	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
