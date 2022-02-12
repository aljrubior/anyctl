package cmd

import (
	"github.com/spf13/cobra"
)

var schedulersCmd = &cobra.Command{
	Use:     "schedulers",
	Aliases: []string{"scheduler", "sch"},
	Short:   "",
}

func init() {
	deploymentsCmd.AddCommand(schedulersCmd)
}
