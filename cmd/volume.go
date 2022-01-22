package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
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
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Volume(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}
