package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describeRuntimeFabricCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Describe a Runtime Fabric instance",
	Run: func(cmd *cobra.Command, args []string) {

		runtimeFabricManager := handlers.NewDefaultRuntimeFabricHandler(ConfigManager, OrganizationRuntimeFabricManager)

		switch len(args) {
		case 1:
			if err := runtimeFabricManager.DescribeFabric(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets organizationPrivateSpaces describe <private-space-name>'")
		}
	},
}

func init() {
	runtimeFabricsCmd.AddCommand(describeRuntimeFabricCmd)
}
