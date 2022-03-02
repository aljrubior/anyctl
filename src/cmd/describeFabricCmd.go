package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describeFabricCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Describe a shared spaces",
	Run: func(cmd *cobra.Command, args []string) {

		fabricHandler := handlers.NewDefaultFabricHandler(ConfigManager, FabricManager, AccountManager)

		switch len(args) {
		case 1:
			if err := fabricHandler.DescribeFabric(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl sharedspaces describe [fabric-name|fabric-id]'")
		}
	},
}

func init() {
	fabricsCmd.AddCommand(describeFabricCmd)
}
