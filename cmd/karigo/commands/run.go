package commands

import (
	"fmt"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/drivers/memory"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			port = 8080
		)

		fmt.Printf("Loading...")
		fmt.Printf(" done.\n")
		fmt.Printf("Listening on port %d...\n", port)

		// Server
		server := &karigo.Server{
			Nodes: map[string]*karigo.Node{},
		}

		src := &memory.Source{}
		_ = src.Reset()
		node := karigo.NewNode(&memory.Journal{}, src)
		node.Name = "test"
		node.Domains = []string{"localhost", "127.0.0.1"}
		for _, domain := range node.Domains {
			server.Nodes[domain] = node
		}

		server.Run()
	},
}
