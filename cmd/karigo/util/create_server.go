package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := karigo.NewServer(config)

	src := &memory.Source{}
	_ = src.Reset()

	// Add cluster control node
	ctlNode := karigo.NewNode(&memory.Journal{}, src)
	ctlNode.Name = "main_node"

	// Register node
	server.Nodes[server.Host] = ctlNode

	for _, host := range server.OtherHosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
