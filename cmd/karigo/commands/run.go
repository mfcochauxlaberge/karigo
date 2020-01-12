package commands

import (
	"fmt"

	"github.com/mfcochauxlaberge/karigo/cmd/karigo/util"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Loading...")
		fmt.Printf(" done.\n")
		fmt.Printf("Listening on port %d...\n", *port)

		server := util.CreateServer()

		server.Run(*port)
	},
}

var (
	port *uint
)

func init() {
	port = cmdRun.Flags().UintP("port", "p", 6280, "port number")
}
