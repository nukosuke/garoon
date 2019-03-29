package cmd

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	//	"time"

	"github.com/otoyo/garoon"
	"github.com/spf13/cobra"
)

var event = &cobra.Command{
	Use:   "event",
	Short: "予定の取得コマンド",
}

var eventList = &cobra.Command{
	Use:   "ls",
	Short: "予定リストを取得します。",

	// TODO: options
	// -r 2019-03-01:2019-03-31
	Run: func(cmd *cobra.Command, args []string) {
		v := url.Values{}
		v.Add("fields", "id,start,end,subject")

		// TODO: paging
		pager, err := client.SearchEvents(v)
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}

		for _, event := range pager.Events {
			fmt.Printf("%d\t[%s ~ %s]\t%s\n", event.ID, event.Start.DateTime.Format("2006/01/02 15:04"), event.End.DateTime.Format("2006/01/02 15:04"), event.Subject)
		}
	},
}

var eventInfo = &cobra.Command{
	Use:   "info",
	Short: "予定の詳細を取得します。",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Println("エラー: 不正なIDです。")
			os.Exit(1)
		}

		event, err := client.FindEvent(id)
		if err != nil {
			fmt.Println("エラー: 予定の取得に失敗しました。")
			os.Exit(1)
		}

		// FIXME: use template
		fmt.Printf(`Subject: %s
Start: %s
End: %s
Attendees: %s
Facilities: %s
Event-Type: %s
ID: %d
Repeat-ID: %s

%s
`,
			event.Subject,
			event.Start.DateTime.Format("Mon Jan _2 15:04:05 2006"),
			event.End.DateTime.Format("Mon Jan _2 15:04:05 2006"),
			joinAttendees(event.Attendees, ", "),
			joinFacilities(event.Facilities, ", "),
			event.EventType,
			event.ID,
			noneIfEmpty(event.RepeatID),
			event.Notes,
		)
	},
}

func init() {
	event.AddCommand(eventList)
	event.AddCommand(eventInfo)
}

// TODO: move to utility.go

func joinAttendees(attendees []garoon.Attendee, separator string) string {
	names := []string{}
	for _, a := range attendees {
		names = append(names, a.Name)
	}
	return strings.Join(names, separator)
}

func joinFacilities(facilities []garoon.Facility, separator string) string {
	names := []string{}
	for _, f := range facilities {
		names = append(names, f.Name)
	}

	if len(names) == 0 {
		return "None"
	}
	return strings.Join(names, separator)
}

func noneIfEmpty(s string) string {
	if s == "" {
		return "None"
	}
	return s
}
