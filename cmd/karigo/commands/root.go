package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

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
