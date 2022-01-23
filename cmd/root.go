package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"spotify-go-cli/service"
)

const (
	configType = "yml"
	configName = ".spotify-go"
	configFile = configName + "." + configType
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use: "spotify-go-client",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

// Execute https://github.com/spf13/cobra/blob/master/user_guide.md#create-rootcmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file ($HOME/.spotify-go.yml)")
}

// initConfig https://github.com/spf13/cobra/blob/master/user_guide.md#create-rootcmd
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			return
		}
		viper.AddConfigPath(home)
		viper.SetConfigType(configType)
		viper.SetConfigName(configName)
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		createConfigFile()
	} else {
		// check if token expired
		if err = service.RefreshToken(); err != nil {
			fmt.Printf("Error while refreshing token. %v\n", err)
			return
		}
	}
}

// createConfigFile creates a config file if there is no
func createConfigFile() {
	config := []byte(`
access_token:
expires_in:
refresh_token:
scope:
token_type:
client_id:
secret:
`)
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err.Error())
	}

	file, err := os.Create(home + "\\" + configFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = file.Write(config)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}
