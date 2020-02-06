package commands

import (
	"github.com/spf13/cobra"
)

var cmdCheck = &cobra.Command{
	Use:   "new",
	Short: "Check instance",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
