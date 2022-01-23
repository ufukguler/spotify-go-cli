package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
	"strconv"
)

func init() {
	rootCmd.AddCommand(volumeCmd)
}

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "Set Playback Volume",
	Long: `Set the volume for the user’s current playback device
The volume to set. Must be a value from 0 to 100 inclusive.`,
	Example: "spotify-go-client volume 50",
	Aliases: []string{"vol"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of volume.")
			return
		}
		parseInt, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Given volume value is not a number.")
			return
		}
		if parseInt < 0 || parseInt > 100 {
			fmt.Println("Invalid volume value. (should be between 0 to 100)")
			return
		}
		if err := service.Volume(args[0]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
