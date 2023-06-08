package main

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use: "thomasscothamilton",
		Run: func(cmd *cobra.Command, args []string) {
			// code you wish you had...
		},
	}
)

func init() {
	// initialize
}

// Execute root command
func Execute() error {
	return rootCmd.Execute()
}
