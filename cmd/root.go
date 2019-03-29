package cmd

import (
	"fmt"
	"os"

	"github.com/otoyo/garoon"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var client *garoon.Client

var root = &cobra.Command{
	Use:   "garoon",
	Short: "Cybozu Garoonコマンド",
	Long: `
garoon v0.0.0 --- Cybozu Garoonのコマンドラインツール

        バグレポートはこちらにお願いします。
        https://github.com/nukosuke/garoon`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		c, err := garoon.NewClient(
			viper.GetString("subdomain"),
			viper.GetString("username"),
			viper.GetString("password"))

		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}
		client = c
	},
}

// サブコマンド追加
func init() {
	root.AddCommand(event)
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println("エラー: ", err)
		os.Exit(1)
	}
}
