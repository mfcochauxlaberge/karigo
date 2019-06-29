package commands

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mfcochauxlaberge/karigo"

	"github.com/spf13/cobra"
)

var cmdRun = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			port = 8080
		)

		fmt.Printf("Loading...")
		time.Sleep(time.Second)
		fmt.Printf(" done.\n")
		fmt.Printf("Listening on port %d...\n", port)

		// Server
		server := &karigo.Server{}

		err := http.ListenAndServe(":"+strconv.Itoa(port), server)
		if err != http.ErrServerClosed {
			panic(err)
		}
	},
}
