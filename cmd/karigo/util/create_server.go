package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := karigo.NewServer()

	src := &memory.Source{}
	_ = src.Reset()

	// Add cluster control node
	ctlNode := karigo.NewNode(&memory.Journal{}, src)
	ctlNode.Name = "main_node"

	// Register node
	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
