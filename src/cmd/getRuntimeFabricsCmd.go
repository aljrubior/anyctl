package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getRuntimeFabrics = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Lists all the deployment targets of type Runtime Fabric available to deploy applications in the current environment.",
	Run: func(cmd *cobra.Command, args []string) {

		runtimeFabricHandler := handlers.NewDefaultRuntimeFabricHandler(ConfigManager, OrganizationRutimeFabricManager)

		switch len(args) {
		case 0:
			if err := runtimeFabricHandler.GetFabrics(); err != nil {
				errors.Catch(err).Println()
			}
		case 1:
			if err := runtimeFabricHandler.FindRuntimeFabricContainsName(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets runtimefabrics get'")
		}
	},
}

func init() {
	runtimeFabricsCmd.AddCommand(getRuntimeFabrics)
}
