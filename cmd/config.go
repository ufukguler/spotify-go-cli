package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show current configuration",
	Run: func(cmd *cobra.Command, args []string) {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err.Error())
		}
		file, err := os.ReadFile(home + "\\.spotify-go.env")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(file))
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
