package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(repeatCmd)
}

var repeatCmd = &cobra.Command{
	Use:       "repeat",
	Short:     "Set Repeat Mode",
	Long:      "Set the repeat mode for the current playing song. Options are track, context, and off.",
	ValidArgs: []string{"track", "context", "off"},
	Example:   "spotify-go-client repeat track",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Repeat(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}
