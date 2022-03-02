package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var describeOrganizationPrivateSpaceCmd = &cobra.Command{
	Use:     "describe",
	Aliases: []string{"desc"},
	Short:   "Describe a private spaces instance",
	Run: func(cmd *cobra.Command, args []string) {

		privateSpaceHandler := handlers.NewDefaultOrganizationPrivateSpaceHandler(ConfigManager, OrganizationPrivateSpaceManager)

		switch len(args) {
		case 1:
			if err := privateSpaceHandler.DescribePrivateSpace(args[0]); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager privatespaces describe <private-space-name>'")
		}
	},
}

func init() {
	organizationPrivateSpacesCmd.AddCommand(describeOrganizationPrivateSpaceCmd)
}
