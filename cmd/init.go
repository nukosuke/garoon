package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// 設定変数
var configFile string
var subdomain string
var username string
var password string

func init() {
	cobra.OnInitialize(initConfig)

	root.PersistentFlags().StringVar(&configFile, "config", "", "設定ファイルのパス (デフォルトは $HOME/.config/garoon/config.yml)")
	root.PersistentFlags().StringVarP(&subdomain, "subdomain", "d", "", "Garoonのサブドメイン")
	root.PersistentFlags().StringVarP(&username, "username", "u", "", "ログインユーザ名")
	root.PersistentFlags().StringVarP(&password, "password", "p", "", "ログインユーザのパスワード")

	viper.BindPFlag("subdomain", root.PersistentFlags().Lookup("subdomain"))
	viper.BindPFlag("username", root.PersistentFlags().Lookup("username"))
	viper.BindPFlag("password", root.PersistentFlags().Lookup("password"))
}

func initConfig() {
	if username != "" && password != "" && subdomain != "" {
		return
	}

	// TODO: macOS key chain option

	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		configPath, err := homedir.Expand("~/.config/garoon/")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(`エラー: 設定ファイルの読み込みに失敗しました。`)
		os.Exit(1)
	}
}
