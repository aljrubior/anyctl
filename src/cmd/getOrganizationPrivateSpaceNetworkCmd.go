package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getOrganizationPrivateSpaceNetworkCmd = &cobra.Command{
	Use:     "get-network",
	Aliases: []string{"get-nw"},
	Short:   "Lists all the nodes in the specified runtime fabric.",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultOrganizationPrivateSpaceHandler(ConfigManager, OrganizationPrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := privateSpaceHandler.GetNetwork(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets runtimefabrics nodes get <runtime-fabric-name>'")
		}
	},
}

func init() {
	organizationPrivateSpacesCmd.AddCommand(getOrganizationPrivateSpaceNetworkCmd)
}
