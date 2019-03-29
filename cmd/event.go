package cmd

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"text/template"
	//	"time"

	"github.com/otoyo/garoon"
	"github.com/spf13/cobra"
)

// TODO: --no-template option
var eventListViewColumns = []string{
	"id",
	"start",
	"end",
	"subject",
}

const (
	eventListViewRowTmpl = "{{.ID}}\t{{.Start.DateTime}}\t{{.End.DateTime}}\t{{.Subject}}\n"

	eventInfoViewTmpl = `Subject: {{.Subject}}
Start: {{.Start.DateTime}}
End: {{.End.DateTime}}
Attendees: {{joinAttendees .Attendees ", "}}
Facilities: {{joinFacilities .Facilities ", "}}
Event-Type: {{.EventType}}
ID: {{.ID}}
Repeat-ID: {{noneIfEmpty .RepeatID}}

{{.Notes}}
`
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
		v.Add("fields", strings.Join(eventListViewColumns, ","))

		// TODO: paging
		pager, err := client.SearchEvents(v)
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}

		t := template.
			Must(template.
				New("row").
				Parse(eventListViewRowTmpl))
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}

		var buf bytes.Buffer
		for _, event := range pager.Events {
			if err = t.Execute(&buf, event); err != nil {
				fmt.Println("エラー: ", err)
				os.Exit(1)
			}
		}
		fmt.Print(buf.String())
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

		t := template.
			Must(template.
				New("info").
				Funcs(template.FuncMap{
					"joinAttendees":  joinAttendees,
					"joinFacilities": joinFacilities,
					"noneIfEmpty":    noneIfEmpty,
				}).
				Parse(eventInfoViewTmpl))
		if err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}

		var buf bytes.Buffer
		if err = t.Execute(&buf, event); err != nil {
			fmt.Println("エラー: ", err)
			os.Exit(1)
		}

		fmt.Print(buf.String())

		// FIXME: use template
		/*
			fmt.Printf(eventInfoViewTmpl,
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
		*/
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
