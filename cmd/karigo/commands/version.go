package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	pretty *bool
)

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: "Show the version number",
	// Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if *pretty {
			fmt.Printf("Karigo v0.0.0\n")
		} else {
			fmt.Printf("v0.0.0\n")
		}
	},
}

func init() {
	pretty = cmdVersion.Flags().BoolP("pretty", "p", false, "show prettier output")
}
