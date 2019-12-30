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
		fmt.Printf("Loading...")
		fmt.Printf(" done.\n")
		fmt.Printf("Listening on port %d...\n", *port)

		// Server
		server := karigo.NewServer()

		var src karigo.Source
		src = &memory.Source{}
		_ = src.Reset()

		// Add cluster control node
		ctlNode := &karigo.Node{
			Name: "0_ctl_node_1",
		}
		ctlNode.AddSource("main", src)
		ctlNode.RegisterJournal(&memory.Journal{})

		// Add empty cluster schema to node
		sc := karigo.ClusterSchema()
		ops := karigo.SchemaToOps(sc)
		err := ctlNode.apply(ops)
		if err != nil {
			panic(err)
		}

		// Register node
		server.Nodes[ctlNode.Name] = ctlNode

		server.Run(*port)
	},
}

var (
	port *uint
)

func init() {
	port = cmdRun.Flags().UintP("port", "p", 6280, "port number")
}
