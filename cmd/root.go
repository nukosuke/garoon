package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "garoon",
	Short: "Cybozu Garoonコマンド",
	Long: `Cybozu Garoonのコマンドラインツールです。

        バグレポートはこちらにお願いします。
        https://github.com/nukosuke/garoon`,
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
