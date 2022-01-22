package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
)

var devicesCmd = &cobra.Command{
	Use:     "devices",
	Short:   "Get Available Devices",
	Long:    "Get information about a user’s available devices.",
	Example: "spotify-go-client devices",
	Run: func(cmd *cobra.Command, args []string) {
		if err := service.Devices(); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(devicesCmd)
}
