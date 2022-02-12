package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getTargetsStandaloneDetailsCmd = &cobra.Command{
	Use:     "get-details",
	Aliases: []string{"get-detail"},
	Short:   "Display the target details",
	Run: func(cmd *cobra.Command, args []string) {

		targetHandler := handlers.NewDefaultTargetHandler(*ConfigManager, *TargetManager)

		switch len(args) {
		case 1:
			if err := targetHandler.GetDetails(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager targets --help'")
		}
	},
}

func init() {
	targetsCmd.AddCommand(getTargetsStandaloneDetailsCmd)
}
