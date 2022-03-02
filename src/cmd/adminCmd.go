package cmd

import "github.com/spf13/cobra"

var adminCmd = &cobra.Command{
	Use:     "admin",
	Aliases: []string{"adm"},
	Short:   "Administrator resources",
	Run: func(cmd *cobra.Command, args []string) {

		println("Error: Unsupported option. Try with 'anyctl admin --help'")
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
