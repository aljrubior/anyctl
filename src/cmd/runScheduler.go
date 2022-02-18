package cmd

import (
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var runSchedulerCmd = &cobra.Command{
	Use:   "run",
	Short: "Run scheduler",
	Run: func(cmd *cobra.Command, args []string) {

		schedulerHandler := handlers.NewDefaultSchedulerHandler(ConfigManager, DeploymentManager, SchedulerManager)

		switch len(args) {
		case 2:
			if err := schedulerHandler.RunScheduler(args[0], args[1]); err != nil {
				errors.Catch(err).Println()
				return
			}
			Console.LogRunSchedulerSuccess(args[1])
		default:
			Console.LogInvalidParameters()
		}
	},
}

func init() {
	schedulersCmd.AddCommand(runSchedulerCmd)
}
