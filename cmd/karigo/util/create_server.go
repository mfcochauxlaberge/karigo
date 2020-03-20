package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
	"github.com/mfcochauxlaberge/karigo/drivers/psql"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := karigo.NewServer()
	server.Config = config

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

	_ = src.Reset(karigo.FirstSchema())

	// Add cluster control node
	ctlNode := karigo.NewNode(jrnl, src)
	ctlNode.Name = "main_node"

	ctlNode.Hosts = config.Hosts
	ctlNode.Journal = config.Journal
	ctlNode.Sources = config.Sources

	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
