package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(pauseCmd)
}

var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "Pause Playback",
	Long:    "Pause playback on the user's account.",
	Example: "spotify-go-client pause",
	Aliases: []string{"stop"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.PauseResume(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
