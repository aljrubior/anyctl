package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getTargetsStandaloneAddressesCmd = &cobra.Command{
	Use:     "get-addresses",
	Aliases: []string{"get-addresse", "get-addr"},
	Short:   "Display the target details",
	Run: func(cmd *cobra.Command, args []string) {

		targetHandler := handlers.NewDefaultTargetHandler(*ConfigManager, *TargetManager)

		switch len(args) {
		case 1:
			if err := targetHandler.GetAddresses(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager targets --help'")
		}
	},
}

func init() {
	targetsCmd.AddCommand(getTargetsStandaloneAddressesCmd)
}
