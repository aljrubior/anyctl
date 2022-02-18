package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getUsageAll *bool

var getOrganizationUsageCmd = &cobra.Command{
	Use:     "get-usage",
	Aliases: []string{"gu"},
	Short:   " organization",
	Run: func(cmd *cobra.Command, args []string) {

		usageHandler := handlers.NewDefaultOrganizationHandler(AccountManager, ConfigManager)

		switch len(args) {
		case 0:

			if *getUsageAll {
				if err := usageHandler.GetAllOrganizationsUsage(); err != nil {
					errors.Catch(err).Println()
				}

				return
			}

			if err := usageHandler.GetCurrentOrganizationUsage(); err != nil {
				errors.Catch(err).Println()
			}
		case 1:
			if err := usageHandler.GetSingleOrganizationUsage(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl organizations get-usage'")
		}
	},
}

func init() {
	organizationsCmd.AddCommand(getOrganizationUsageCmd)
	getUsageAll = getOrganizationUsageCmd.Flags().Bool("all", false, "Include Sub Organizations")
}
