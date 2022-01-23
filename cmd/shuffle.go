package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(shuffleCmd)
}

var shuffleCmd = &cobra.Command{
	Use:       "shuffle",
	Short:     "Toggle Playback Shuffle",
	Long:      "Toggle shuffle on or off for user’s playback.",
	Example:   "spotify-go-client shuffle on",
	ValidArgs: []string{"on", "off"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of Shuffle.")
			return
		}
		if err := service.Shuffle(args[0]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
