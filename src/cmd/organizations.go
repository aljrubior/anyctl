package cmd

import (
	"github.com/spf13/cobra"
)

var organizationsCmd = &cobra.Command{
	Use:     "organizations",
	Aliases: []string{"organization", "org"},
	Short:   "Organization resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl organizations --help'")
	},
}

func init() {
	adminCmd.AddCommand(organizationsCmd)
}
