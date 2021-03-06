package util

import (
	"github.com/mfcochauxlaberge/karigo"
)

func CreateServer(config karigo.Config) *karigo.Server {
	// Server
	server := karigo.NewServer()
	server.Config = config

	// Add cluster control node
	ctlNode := karigo.NewNode(config)
	ctlNode.Name = "main_node"

	for _, host := range config.Hosts {
		server.Nodes[host] = ctlNode
	}

	return server
}
