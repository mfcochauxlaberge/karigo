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
		cmdConfig,
		cmdExec,
		cmdNew,
		cmdRun,
		cmdVersion,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		Config, err = util.ReadConfig(*configPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not read configuration file: %s\n", err)
			os.Exit(1)
		}
	},
}

var (
	configPath *string
)

func init() {
	configPath = rootCmd.PersistentFlags().StringP(
		"config", "c",
		"karigo.yml",
		"configuration file",
	)
}
