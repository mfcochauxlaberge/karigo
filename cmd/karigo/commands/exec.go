package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	commit *bool
)

var cmdExec = &cobra.Command{
	Use:   "exec",
	Short: "Execute an operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Operation is valid.\n")

		if *commit {
			fmt.Printf("Operation committed.\n")
		}
	},
}

func init() {
	commit = cmdExec.Flags().BoolP("commit", "", false, "commit the operation")
}
