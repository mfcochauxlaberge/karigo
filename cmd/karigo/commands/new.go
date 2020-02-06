package commands

import (
	"github.com/spf13/cobra"
)

var cmdNew = &cobra.Command{
	Use:   "new",
	Short: "Create new instance",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
	},
}
