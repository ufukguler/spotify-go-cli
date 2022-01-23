package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"spotify-go-cli/service"
	"strings"
)

func init() {
	rootCmd.AddCommand(searchCmd)
}

var searchCmd = &cobra.Command{
	Use:     "search",
	Short:   "Search for Item",
	Long:    "Get Spotify catalog information about tracks that match a keyword string. Flags are optional.",
	Example: "spotify-go-client search [query] --limit=5",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Invalid usage of search.")
			return
		}

		input := strings.Join(args, " ")
		if err := service.Search(input); err != nil {
			fmt.Println(err)
		}
	},
}
