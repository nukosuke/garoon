package cmd

import (
	"fmt"
	"net/url"
	"os"

	_ "github.com/otoyo/garoon"
	"github.com/spf13/cobra"
)

var event = &cobra.Command{
	Use:   "event",
	Short: "予定を取得します。",
	Run: func(cmd *cobra.Command, args []string) {
		v := url.Values{}

		// TODO: paging
		pager, err := client.SearchEvents(v)
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}
		fmt.Println(pager.Events)
	},
}
