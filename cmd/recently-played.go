package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

var recentlyPlayedCmd = &cobra.Command{
	Use:     "recent",
	Short:   "Get Recently Played Tracks",
	Long:    "Get tracks from the current user's recently played tracks.",
	Example: "spotify-go-client recent",
	Aliases: []string{"rp"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.RecentlyPlayed(); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(recentlyPlayedCmd)
}
