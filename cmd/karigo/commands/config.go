package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cmdConfig = &cobra.Command{
	Use:   "config",
	Short: "Show the configuration (JSON)",
	Run: func(cmd *cobra.Command, args []string) {
		content, err := json.MarshalIndent(Config, "", "\t")
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Println(string(content))
	},
}
