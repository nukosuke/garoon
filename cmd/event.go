package cmd

import (
	"net/url"

	_ "github.com/otoyo/garoon"
	"github.com/spf13/cobra"
)

var event = &cobra.Command{
	Use:   "event",
	Short: "予定を取得します。",
	Run: func(cmd *cobra.Command, args []string) {
		v := url.Values{}
		// todo
		client.SearchEvents(v)
	},
}
