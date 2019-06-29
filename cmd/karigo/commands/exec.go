package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdExec = &cobra.Command{
	Use:   "exec",
	Short: "Execute an operation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Operation committed.\n")
	},
}
