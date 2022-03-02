package cmd

import (
	"github.com/spf13/cobra"
)

var assetsCmd = &cobra.Command{
	Use:     "assets",
	Aliases: []string{"asset"},
	Short:   "Asset resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl assets --help'")
	},
}

func init() {
	rootCmd.AddCommand(assetsCmd)
}
