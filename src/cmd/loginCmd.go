package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var username string
var password string

var loginCmd = &cobra.Command{
	Use:     "login",
	Aliases: []string{"lo"},
	Short:   "Log into Anypoint platform",
	Run: func(cmd *cobra.Command, args []string) {

		loginHandler := handlers.NewDefaultLoginHandler(AccountManager, ConfigManager)

		if err := loginHandler.Login(username, password); err != nil {
			Console.LogError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVarP(&username, "username", "U", "", "Anypoint username (required)")
	loginCmd.MarkFlagRequired("username")

	loginCmd.Flags().StringVarP(&password, "password", "P", "", "Anypoint password (required)")
	loginCmd.MarkFlagRequired("password")
}
