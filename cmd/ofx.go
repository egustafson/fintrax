package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ofxCmd = &cobra.Command{
	Use:  "ofx",
	RunE: doOFX,
}

func init() {
	rootCmd.AddCommand(ofxCmd)
}

func doOFX(cmd *cobra.Command, args []string) error {

	fmt.Println("wip...")

	return nil
}
