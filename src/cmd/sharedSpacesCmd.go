package cmd

import "github.com/spf13/cobra"

var sharedSpacesCmd = &cobra.Command{
	Use:     "sharedspaces",
	Aliases: []string{"sharedspace", "ss"},
	Short:   "Shared Space resource",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl sharedspaces --help'")
	},
}

func init() {
	adminCmd.AddCommand(sharedSpacesCmd)
}
