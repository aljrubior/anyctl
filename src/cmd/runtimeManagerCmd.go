package cmd

import "github.com/spf13/cobra"

var runtimeManagerCmd = &cobra.Command{
	Use:     "runtimemanager",
	Aliases: []string{"rm"},
	Short:   "Runtime Manager Interface",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl runtimemanager --help'")
	},
}

func init() {
	rootCmd.AddCommand(runtimeManagerCmd)
}
