package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describeSharedSpaceCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"d"},
	Short:   "Describe a shared spaces",
	Run: func(cmd *cobra.Command, args []string) {

		sharedSpaceHandler := handlers.NewDefaultSharedSpaceHandler(*ConfigManager, *SharedSpaceManager, *PrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := sharedSpaceHandler.DescribeSharedSpace(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl sharedspaces describe <shared-space-name>'")
		}
	},
}

func init() {
	sharedSpacesCmd.AddCommand(describeSharedSpaceCmd)
}
