package time

import "time"

const (
	DateYearFormat       = "2006"
	DateMonthFormat      = "2006-01"
	DateDayFormat        = "2006-01-02"
	DateTimeHourFormat   = "2006-01-02 15"
	DateTimeMinuteFormat = "2006-01-02 15:04"
	DateTimeSecondFormat = "2006-01-02 15:04:05"
	DateTimeMilliFormat  = "2006-01-02 15:04:05.999"
	DateTimeMicroFormat  = "2006-01-02 15:04:05.999999"
)

// Formats default time formats will be parsed as
var Formats = []string{
	"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2",
	"15:4:5", "15:4", "15",
	"15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05Z0700", "2006-01-02T15:04:05Z07",
	"2006.1.2", "2006.1.2 15:04:05", "2006.01.02", "2006.01.02 15:04:05", "2006.01.02 15:04:05.999999999",
	"1/2/2006", "1/2/2006 15:4:5", "2006/01/02", "20060102", "2006/01/02 15:04:05",
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
	time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

func formatTimeToList(t time.Time) []int {
	hour, minute, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, minute, hour, day, int(month), year}
}

func ParseToYear(t time.Time) string {
	return t.Format(DateYearFormat)
}

func ParseToMonth(t time.Time) string {
	return t.Format(DateMonthFormat)
}

func ParseToDay(t time.Time) string {
	return t.Format(DateDayFormat)
}

func ParseToHour(t time.Time) string {
	return t.Format(DateTimeHourFormat)
}

func ParseToMinute(t time.Time) string {
	return t.Format(DateTimeMinuteFormat)
}

func ParseToSecond(t time.Time) string {
	return t.Format(DateTimeSecondFormat)
}

func ParseToMilli(t time.Time) string {
	return t.Format(DateTimeMilliFormat)
}

func ParseToMicro(t time.Time) string {
	return t.Format(DateTimeMicroFormat)
}
