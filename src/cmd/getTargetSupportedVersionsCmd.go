package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getTargetsSupportedVersions = &cobra.Command{
	Use:     "get-supported-versions",
	Aliases: []string{"get-supported-version"},
	Short:   "Retrieve the list of supported Mule runtimes",
	Run: func(cmd *cobra.Command, args []string) {

		targetHandler := handlers.NewDefaultTargetHandler(ConfigManager, TargetManager)

		switch len(args) {
		case 1:
			if err := targetHandler.GetSupportedRuntimes(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager targets get'")
		}
	},
}

func init() {
	targetsCmd.AddCommand(getTargetsSupportedVersions)
}
