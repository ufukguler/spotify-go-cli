package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(prev)
}

var prev = &cobra.Command{
	Use:     "previous",
	Short:   "Skip To Previous",
	Long:    "Skips to previous track in the user’s queue.",
	Example: "spotify-go-client previous",
	Aliases: []string{"prev"},
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Previous(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
