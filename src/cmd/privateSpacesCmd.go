package cmd

import "github.com/spf13/cobra"

var privateSpacesCmd = &cobra.Command{
	Use:     "privatespaces",
	Aliases: []string{"privatespaces", "ps"},
	Short:   "Private Space resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl targets sharedspaces --help'")
	},
}

func init() {
	adminCmd.AddCommand(privateSpacesCmd)
}
