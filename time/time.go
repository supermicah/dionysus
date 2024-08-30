package time

import "time"

// WeekStartDay set week start day, default is sunday
var WeekStartDay = time.Sunday

// Config configuration for now package
type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

// DefaultConfig default config
var DefaultConfig *Config

// With New initialize Now based on configuration
func (config *Config) With(t time.Time) *Now {
	return &Now{Time: t, Config: config}
}

// Parse string to time based on configuration
func (config *Config) Parse(strValues ...string) (time.Time, error) {
	if config.TimeLocation == nil {
		return config.With(time.Now()).Parse(strValues...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).Parse(strValues...)
	}
}

// MustParse must parse string to time or will panic
func (config *Config) MustParse(strValues ...string) time.Time {
	if config.TimeLocation == nil {
		return config.With(time.Now()).MustParse(strValues...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).MustParse(strValues...)
	}
}

// Now struct
type Now struct {
	time.Time
	*Config
}

// With initialize Now with time
func With(t time.Time) *Now {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
			TimeFormats:  Formats,
		}
	}

	return &Now{Time: t, Config: config}
}

// New initialize Now with time
func New(t time.Time) *Now {
	return With(t)
}

// BeginningOfMinute beginning of minute
func BeginningOfMinute() time.Time {
	return With(time.Now()).BeginningOfMinute()
}

// BeginningOfHour beginning of hour
func BeginningOfHour() time.Time {
	return With(time.Now()).BeginningOfHour()
}

// BeginningOfDay beginning of day
func BeginningOfDay() time.Time {
	return With(time.Now()).BeginningOfDay()
}

// BeginningOfWeek beginning of week
func BeginningOfWeek() time.Time {
	return With(time.Now()).BeginningOfWeek()
}

// BeginningOfMonth beginning of month
func BeginningOfMonth() time.Time {
	return With(time.Now()).BeginningOfMonth()
}

// BeginningOfQuarter beginning of quarter
func BeginningOfQuarter() time.Time {
	return With(time.Now()).BeginningOfQuarter()
}

// BeginningOfYear beginning of year
func BeginningOfYear() time.Time {
	return With(time.Now()).BeginningOfYear()
}

// EndOfMinute end of minute
func EndOfMinute() time.Time {
	return With(time.Now()).EndOfMinute()
}

// EndOfHour end of hour
func EndOfHour() time.Time {
	return With(time.Now()).EndOfHour()
}

// EndOfDay end of day
func EndOfDay() time.Time {
	return With(time.Now()).EndOfDay()
}

// EndOfWeek end of week
func EndOfWeek() time.Time {
	return With(time.Now()).EndOfWeek()
}

// EndOfMonth end of month
func EndOfMonth() time.Time {
	return With(time.Now()).EndOfMonth()
}

// EndOfQuarter end of quarter
func EndOfQuarter() time.Time {
	return With(time.Now()).EndOfQuarter()
}

// EndOfYear end of year
func EndOfYear() time.Time {
	return With(time.Now()).EndOfYear()
}

// Monday monday

func Monday(strValues ...string) time.Time {
	return With(time.Now()).Monday(strValues...)
}

// Sunday get sunday
func Sunday(strValues ...string) time.Time {
	return With(time.Now()).Sunday(strValues...)
}

// EndOfSunday end of sunday
func EndOfSunday() time.Time {
	return With(time.Now()).EndOfSunday()
}

// Quarter returns the yearly quarter
func Quarter() uint {
	return With(time.Now()).Quarter()
}

// Parse string to time
func Parse(strValues ...string) (time.Time, error) {
	return With(time.Now()).Parse(strValues...)
}

// ParseInLocation parse string to time in location
func ParseInLocation(loc *time.Location, strValues ...string) (time.Time, error) {
	return With(time.Now().In(loc)).Parse(strValues...)
}

// MustParse must parse string to time or will panic
func MustParse(strValues ...string) time.Time {
	return With(time.Now()).MustParse(strValues...)
}

// MustParseInLocation must parse string to time in location or will panic
func MustParseInLocation(loc *time.Location, strValues ...string) time.Time {
	return With(time.Now().In(loc)).MustParse(strValues...)
}

// Between check now between the start, end time or not
func Between(time1, time2 string) bool {
	return With(time.Now()).Between(time1, time2)
}
