package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(queueCmd)
}

var queueCmd = &cobra.Command{
	Use:     "queue",
	Short:   "Add Item to Playback Queue",
	Long:    "Add an item to the end of the user's current playback queue.",
	Example: "spotify-go-client queue [uri]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of queue.")
			return
		}

		if err := service.Queue(args[0]); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Added item to playback queue.")
	},
}
