package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var getTargetsCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Lists all the deployment targets available to deploy applications in the current environment.",
	Run: func(cmd *cobra.Command, args []string) {

		targetHandler := handlers.NewDefaultTargetHandler(ConfigManager, TargetManager)

		switch len(args) {
		case 0:
			if err := targetHandler.GetTargets(); err != nil {
				errors.Catch(err).Println()
			}
		case 1:
			if err := targetHandler.FindTargetsContainsName(args[0]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager targets get'")
		}
	},
}

func init() {
	targetsCmd.AddCommand(getTargetsCmd)
}
