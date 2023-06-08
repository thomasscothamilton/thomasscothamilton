package cmd

import (
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "initialize",
	Run: func(cmd *cobra.Command, args []string) {
		// code you wish you had...
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
