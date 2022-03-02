package cmd

import "github.com/spf13/cobra"

var organizationPrivateSpacesCmd = &cobra.Command{
	Use:     "privatespaces",
	Aliases: []string{"privatespaces", "ps"},
	Short:   "Private Space resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl targets privatespaces --help'")
	},
}

func init() {
	runtimeManagerCmd.AddCommand(organizationPrivateSpacesCmd)
}
