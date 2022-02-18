package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getRuntimeFabricNodeCmd = &cobra.Command{
	Use:     "get-nodes",
	Aliases: []string{"get-nodes", "get-no"},
	Short:   "Lists all the nodes in the specified runtime fabric.",
	Run: func(cmd *cobra.Command, args []string) {

		runtimeFabricHandler := handlers.NewDefaultRuntimeFabricHandler(ConfigManager, OrganizationRutimeFabricManager)

		switch len(args) {
		case 1:
			if err := runtimeFabricHandler.GetRuntimeFabricNodes(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets runtimefabrics nodes get <runtime-fabric-name>'")
		}
	},
}

func init() {
	runtimeFabricsCmd.AddCommand(getRuntimeFabricNodeCmd)
}
