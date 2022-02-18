package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getSharedSpacesCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of shared spaces",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultSharedSpaceHandler(ConfigManager, SharedSpaceManager, PrivateSpaceManager)

		switch len(args) {
		case 0:
			if err := privateSpaceHandler.GetSharedSpaces(); err != nil {
				Console.LogError(err)
			}
		case 1:
			if err := privateSpaceHandler.FindSharedSpaceContainsName(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets sharedspaces get'")
		}
	},
}

func init() {
	sharedSpacesCmd.AddCommand(getSharedSpacesCmd)
}
