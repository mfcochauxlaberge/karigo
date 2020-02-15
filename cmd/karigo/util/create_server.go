package util

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"
)

func CreateServer() *karigo.Server {
	// Server
	server := karigo.NewServer()

	src := &memory.Source{}
	_ = src.Reset()

	// Add cluster control node
	ctlNode := karigo.NewNode(&memory.Journal{}, src)
	ctlNode.Name = "main_node"

	// Register node
	server.Nodes["127.0.0.1"] = ctlNode
	server.Nodes["localhost"] = ctlNode

	return server
}
