package time

import "time"

func formatTimeToList(t time.Time) []int {
	hour, minute, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, minute, hour, day, int(month), year}
}
