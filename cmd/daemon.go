package cmd

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use: "daemon",
	RunE: doDaemon,
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func doDaemon(cmd *cobra.Command, args []string) error {

	var err error = nil // invoke the daemon
	return err
}
