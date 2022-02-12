package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getOrganizationQuotasCmd = &cobra.Command{
	Use:     "get-quotas",
	Aliases: []string{"gq"},
	Short:   "Retrieve current organization quotas",
	Run: func(cmd *cobra.Command, args []string) {

		entitlementHandler := handlers.NewDefaultOrganizationHandler(*AccountManager, *ConfigManager)

		switch len(args) {
		case 0:
			if err := entitlementHandler.GetCurrentOrganizationQuotas(); err != nil {
				Console.LogError(err)
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl organizations get-quotas'")
		}
	},
}

func init() {
	organizationsCmd.AddCommand(getOrganizationQuotasCmd)
}
