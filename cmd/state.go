package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(stateCmd)
}

var stateCmd = &cobra.Command{
	Use:     "state",
	Short:   "Get Playback State",
	Long:    "Pause playback on the user's account.",
	Example: "spotify-go-client state",
	Aliases: []string{"now"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.State(); err != nil {
			fmt.Println(err)
			return
		}
	},
}
