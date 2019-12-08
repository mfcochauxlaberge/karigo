package commands

import (
	"fmt"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/memory"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			port uint = 6280
		)

		fmt.Printf("Loading...")
		fmt.Printf(" done.\n")
		fmt.Printf("Listening on port %d...\n", port)

		// Server
		server := karigo.NewServer()

		src := &memory.Source{}
		_ = src.Reset()
		node := karigo.NewNode(&memory.Journal{}, src)
		node.Name = "test"
		node.Domains = []string{"localhost", "127.0.0.1"}
		for _, domain := range node.Domains {
			server.Nodes[domain] = node
		}

		server.Run(port)
	},
}
