package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:     "start",
	Short:   "Start/Resume Playback",
	Long:    "Start a new context or resume current playback on the user's active device.",
	Example: "spotify-go-client play",
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Start(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
