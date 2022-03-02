package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Display one or many resources.",
	Run: func(cmd *cobra.Command, args []string) {
		println("get")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
