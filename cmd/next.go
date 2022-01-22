package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

var nextCmd = &cobra.Command{
	Use:     "next",
	Short:   "Skip To Next",
	Long:    "Skips to next track in the queue.",
	Example: "spotify-go-client next",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Next(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
