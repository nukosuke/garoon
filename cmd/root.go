package cmd

import (
	"fmt"
	"os"

	"github.com/otoyo/garoon"
	"github.com/spf13/cobra"
)

var client *garoon.Client

var root = &cobra.Command{
	Use:   "garoon",
	Short: "Cybozu Garoonコマンド",
	Long: `Cybozu Garoonのコマンドラインツールです。

        バグレポートはこちらにお願いします。
        https://github.com/nukosuke/garoon`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client, err := garoon.NewClient(subdomain, username, password)
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}
		_ = client
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
