package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describePrivateSpaceCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "List all private spaces",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultPrivateSpaceHandler(*ConfigManager, *PrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := privateSpaceHandler.DescribePrivateSpace(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl privatespaces get <private-space-id>'")
		}
	},
}

func init() {
	privateSpacesCmd.AddCommand(describePrivateSpaceCmd)
}
