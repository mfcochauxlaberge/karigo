package commands

import (
	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/cmd/karigo/util"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		server := util.CreateServer(karigo.Config{
			Port: 6280,
			Host: "127.0.0.1",
			Journal: map[string]string{
				"type": "memory",
			},
			Sources: map[string]map[string]string{
				"main": {
					"type": "memory",
				},
			},
		})

		server.Run()
	},
}
