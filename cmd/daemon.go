package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/egustafson/fintrax/pkg/server"
)

var daemonCmd = &cobra.Command{
	Use: "daemon",
	Run: doDaemon,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func doDaemon(cmd *cobra.Command, args []string) {
	if err := server.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "server failed to start: %v", err)
	}
}
