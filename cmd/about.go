package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use: "about",
	RunE: doAbout,
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}

func doAbout(cmd *cobra.Command, args []string) error {

	fmt.Println("---")
	fmt.Printf("fintrax-version: %s\n", GitSummary)
	fmt.Printf("build-date: %s\n", BuildDate)
	fmt.Println("...")
	
	return nil
}
