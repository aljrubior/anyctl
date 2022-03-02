package cmd

import (
	"github.com/aljrubior/anyctl/handlers"
	"github.com/spf13/cobra"
)

var refreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "Renew the access token.",
	Run: func(cmd *cobra.Command, args []string) {

		configHandler := handlers.NewDefaultConfigHandler(AccountManager, ConfigManager)

		if err := configHandler.RefreshAccessToken(); err != nil {
			Console.LogError(err)
			return
		}

		Console.LogAccessTokenUpdate()
	},
}

func init() {
	configCmd.AddCommand(refreshTokenCmd)
}
