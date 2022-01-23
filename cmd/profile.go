package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(profileCmd)
}

var profileCmd = &cobra.Command{
	Use:     "profile",
	Short:   "Get Current User's Profile",
	Long:    "Get detailed profile information about the current user (including the current user's username).",
	Example: "spotify-go-client profile",
	Aliases: []string{"user", "account"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Profile(); err != nil {
			fmt.Println(err)
		}
	},
}
