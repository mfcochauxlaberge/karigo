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
		})
		server.Run(*port)
	},
}

var (
	port *uint
)

func init() {
	port = cmdRun.Flags().UintP("port", "p", 6280, "port number")
}
