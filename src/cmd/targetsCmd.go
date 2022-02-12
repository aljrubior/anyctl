package cmd

import (
	"github.com/spf13/cobra"
)

var targetsCmd = &cobra.Command{
	Use:     "targets",
	Aliases: []string{"target"},
	Short:   "Deployment target resource.",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl targets --help'")
	},
}

func init() {
	runtimeManagerCmd.AddCommand(targetsCmd)
}
