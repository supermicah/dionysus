package time

import (
	"fmt"
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	fmt.Println(time.Now()) // 2013-11-18 17:51:49.123456789 Mon

	fmt.Println(BeginningOfMinute())  // 2013-11-18 17:51:00 Mon
	fmt.Println(BeginningOfHour())    // 2013-11-18 17:00:00 Mon
	fmt.Println(BeginningOfDay())     // 2013-11-18 00:00:00 Mon
	fmt.Println(BeginningOfWeek())    // 2013-11-17 00:00:00 Sun
	fmt.Println(BeginningOfMonth())   // 2013-11-01 00:00:00 Fri
	fmt.Println(BeginningOfQuarter()) // 2013-10-01 00:00:00 Tue
	fmt.Println(BeginningOfYear())    // 2013-01-01 00:00:00 Tue

	fmt.Println(EndOfMinute())  // 2013-11-18 17:51:59.999999999 Mon
	fmt.Println(EndOfHour())    // 2013-11-18 17:59:59.999999999 Mon
	fmt.Println(EndOfDay())     // 2013-11-18 23:59:59.999999999 Mon
	fmt.Println(EndOfWeek())    // 2013-11-23 23:59:59.999999999 Sat
	fmt.Println(EndOfMonth())   // 2013-11-30 23:59:59.999999999 Sat
	fmt.Println(EndOfQuarter()) // 2013-12-31 23:59:59.999999999 Tue
	fmt.Println(EndOfYear())    // 2013-12-31 23:59:59.999999999 Tue

	WeekStartDay = time.Monday // Set Monday as first day, default is Sunday
	fmt.Println(EndOfWeek())   // 2013-11-24 23:59:59.999999999 Sun
}
