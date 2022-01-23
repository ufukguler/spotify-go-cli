package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(currentSongCmd)
}

var currentSongCmd = &cobra.Command{
	Use:     "current",
	Short:   "Get Currently Playing Track",
	Long:    "Get the object currently being played on the user's Spotify account.",
	Example: "spotify-go-client current",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Current(); err != nil {
			fmt.Println(err)
		}
	},
}
