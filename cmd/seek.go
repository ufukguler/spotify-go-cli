package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
	"strconv"
)

func init() {
	rootCmd.AddCommand(seekCmd)
}

var seekCmd = &cobra.Command{
	Use:     "seek",
	Short:   "Seek To Position",
	Long:    "Seeks to the given position in the user’s currently playing track.",
	Example: "spotify-go-client seek [SECOND]",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of seek.")
			return
		}
		_, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("Given seek value is not a number.")
			return
		}
		sec := args[0] + "000"
		if err := service.Seek(sec); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[OK]")
	},
}
