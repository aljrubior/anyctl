package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var migrateWithName, migrateToTargetName, migrateToEnvironment *string

var migrateDeploymentCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate deployment",
	Run: func(cmd *cobra.Command, args []string) {

		migrationHandler := handlers.NewDefaultDeploymentMigrationHandler(*DeployerManager, *ConfigManager, *DeploymentManager, *TargetManager, *AccountManager)

		switch len(args) {
		case 1:
			if err := migrationHandler.Migrate(args[0], *migrateWithName, *migrateToTargetName, *migrateToEnvironment); err != nil {
				errors.Catch(err).Println()
			}
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	deploymentsCmd.AddCommand(migrateDeploymentCmd)
	migrateWithName = migrateDeploymentCmd.Flags().StringP("with-name", "", "", "Deployment name ")
	migrateToTargetName = migrateDeploymentCmd.Flags().StringP("to-target-name", "", "", "Target destination")
	migrateToEnvironment = migrateDeploymentCmd.Flags().StringP("to-environment-name", "", "", "Environment destination")
}
