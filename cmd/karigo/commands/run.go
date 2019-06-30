package commands

import (
	"fmt"

	"github.com/mfcochauxlaberge/karigo"

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
		server := &karigo.Server{}
		server.Run()
	},
}
