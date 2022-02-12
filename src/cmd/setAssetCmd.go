package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var setAsset *string

var setAssetCmd = &cobra.Command{
	Use:   "set",
	Short: "Update the deployment asset",
	Run: func(cmd *cobra.Command, args []string) {

		deploymentHandler := handlers.NewDeploymentHandler(*DeploymentManager, *ConfigManager, *TargetManager)

		switch len(args) {
		case 1:
			if err := deploymentHandler.SetDeploymentAsset(args[0], *setAsset); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(setAssetCmd)
	setAsset = setAssetCmd.Flags().StringP("asset", "", "", "Artifact")
}
