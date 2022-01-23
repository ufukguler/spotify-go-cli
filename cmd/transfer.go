package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

func init() {
	rootCmd.AddCommand(transferCmd)
}

var transferCmd = &cobra.Command{
	Use:     "transfer",
	Short:   "Transfer Playback",
	Long:    "Transfer playback to a new device and determine if it should start playing.",
	Example: "spotify-go-client transfer [DEVICE-ID]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of transfer.")
			return
		}
		if err := service.Transfer(args[0]); err != nil {
			fmt.Println(err)
		}
	},
}
