package cmd

import "github.com/spf13/cobra"

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"hist"},
	Short:   "",
}

func init() {
	deploymentsCmd.AddCommand(historyCmd)
}
