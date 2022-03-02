package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var diffHistoryCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff the current deployment spec to previous",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHistoryHandler := handlers.NewDefaultDeploymentHistoryHandler(DeploymentManager, ConfigManager)

		switch len(args) {
		case 2:
			if err := deploymentHistoryHandler.Compare(args[0], args[1]); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager deployments history diff <deployment-name> <spec-version>' ")
		}
	},
}

func init() {
	historyCmd.AddCommand(diffHistoryCmd)
}
