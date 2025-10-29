package cmd

import "github.com/egustafson/fintrax/pkg/config"

// flags for fintrax

var (
	flags = &config.Flags{}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flags.Verbose, "verbose", "v", false,
		"verbose output")
}
