package commands

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Loading...")
		time.Sleep(800 * time.Millisecond)
		fmt.Printf(" done.\n")
		fmt.Printf("Now listening...\n")

		// Server
		server := &karigo.Server{}

		err := http.ListenAndServe(":8080", server)
		if err != http.ErrServerClosed {
			panic(err)
		}
	},
}
