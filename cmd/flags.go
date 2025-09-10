package cmd

// flags for fintrax

var (
	verboseFlag bool = false
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false,
		"verbose outputx")
}
