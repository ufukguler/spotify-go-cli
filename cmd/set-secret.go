package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setSecretKeyCmd = &cobra.Command{
	Use:     "set-secret",
	Short:   "Set Secret Key",
	Example: "spotify-go-client set-secret [SECRET-KEY]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("Invalid Secret Key")
			return
		}
		viper.Set("SECRET", args[0])
		fmt.Printf("Secret Key: %v", viper.GetString("SECRET"))
		err := viper.WriteConfig()

		if err != nil {
			fmt.Printf("%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setSecretKeyCmd)
}
