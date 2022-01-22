package cmd

import (
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login",
	Run: func(cmd *cobra.Command, args []string) {
		service.Login()
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
