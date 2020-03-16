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
	}

	// Source
	var src karigo.Source

	for _, source := range config.Sources {
		switch source["driver"] {
		case "memory":
			src = &memory.Source{}
		default:
		}
	}

	_ = src.Reset()

	// Add cluster control node
	ctlNode := karigo.NewNode(jrnl, src)
	ctlNode.Name = "main_node"

	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
