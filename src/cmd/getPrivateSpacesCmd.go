package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getPrivateSpacesCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "List all private spaces",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultPrivateSpaceHandler(*ConfigManager, *PrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := privateSpaceHandler.GetPrivateSpaces(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl privatespaces get <private-space-name|private-space-id>'")
		}
	},
}

func init() {
	privateSpacesCmd.AddCommand(getPrivateSpacesCmd)
}
