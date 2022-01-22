package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setClientCmd = &cobra.Command{
	Use:     "set-client",
	Short:   "Set Client ID",
	Example: "spotify-go-client set-client [CLIENT-ID]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Invalid Client ID")
			return
		}

		viper.Set("CLIENT_ID", args[0])
		fmt.Printf("Client ID: %v", viper.GetString("CLIENT_ID"))
		err := viper.WriteConfig()

		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setClientCmd)
}
