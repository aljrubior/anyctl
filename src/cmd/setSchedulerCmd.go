package cmd

import (
	"fmt"
	"github.com/aljrubior/anyctl/errors"
	"github.com/aljrubior/anyctl/handlers"
	errors2 "github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
)

var setSchedulerEnabled *string
var setAllSchedulers *bool

var setSchedulerCmd = &cobra.Command{
	Use:   "set",
	Short: "Enable a scheduler in a specified deployment",
	Run: func(cmd *cobra.Command, args []string) {

		schedulerHandler := handlers.NewDefaultSchedulerHandler(ConfigManager, DeploymentManager, SchedulerManager)

		isEnabled := false

		// Parse enabled
		if *setSchedulerEnabled != "" {
			value, err := strconv.ParseBool(*setSchedulerEnabled)

			if err != nil {
				Console.LogError(errors2.New(fmt.Sprintf("Enabled has an invalid value '%s'", *setSchedulerEnabled)))
				return
			}
			isEnabled = value
		}

		switch len(args) {
		case 1:
			if *setAllSchedulers {
				if err := schedulerHandler.EnableSchedulers(args[0], isEnabled); err != nil {
					errors.Catch(err).Println()
				}
				return
			}
		case 2:
			if err := schedulerHandler.EnableScheduler(args[0], args[1], isEnabled); err != nil {
				errors.Catch(err).Println()
			}
		default:
			println("Error: Unsupported option. Try with 'anyctl runtimemanager deployments schedulers set <deployment-name> --enabled <true|false>'")
		}
	},
}

func init() {
	schedulersCmd.AddCommand(setSchedulerCmd)

	setSchedulerEnabled = setSchedulerCmd.Flags().StringP("enabled", "", "", "Enable or disable a scheduler")
	setAllSchedulers = setSchedulerCmd.Flags().BoolP("all", "", false, "Apply the change to all schedulers")
}
