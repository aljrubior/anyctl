package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getOrganizationPrivateSpacesCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Retrieve a list of private spaces",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultOrganizationPrivateSpaceHandler(ConfigManager, OrganizationPrivateSpaceManager)

		switch len(args) {
		case 0:
			if err := privateSpaceHandler.GetPrivateSpaces(); err != nil {
				Console.LogError(err)
			}
		case 1:
			if err := privateSpaceHandler.FindPrivateSpaceContainsName(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl targets privatespaces get'")
		}
	},
}

func init() {
	organizationPrivateSpacesCmd.AddCommand(getOrganizationPrivateSpacesCmd)
}
