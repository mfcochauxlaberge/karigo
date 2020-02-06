package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Execute ...
func Execute() {
	rootCmd.AddCommand(
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

var rootCmd = &cobra.Command{}
