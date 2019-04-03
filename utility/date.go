package utility

import (
	"time"
)

func BeginningOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	return time.Date(y, m, d+1, 0, 0, -1, 0, t.Location())
}

func DateFormat(t time.Time) string {
	return t.Format("2006-01-02T15:04:05-07:00")
}
