package cmd

import (
	"github.com/spf13/cobra"
)

var deploymentsCmd = &cobra.Command{
	Use:     "deployments",
	Aliases: []string{"deploy", "deployment"},
	Short:   "Lists all the deployments in the current environment.",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl deployments --help'")
	},
}

func init() {
	runtimeManagerCmd.AddCommand(deploymentsCmd)
}
