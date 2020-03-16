package commands

import (
	"fmt"
	"os"

	"github.com/mfcochauxlaberge/karigo"
	"github.com/mfcochauxlaberge/karigo/cmd/karigo/util"

	"github.com/spf13/cobra"
)

var (
	Config karigo.Config
)

// Execute ...
func Execute() {
	rootCmd.AddCommand(
		cmdExec,
		cmdRun,
		cmdVersion,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		Config, err = util.ReadConfig("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read configuration file: %s", err)
			os.Exit(1)
		}
	},
}
