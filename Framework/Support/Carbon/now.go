package Carbon

import (
	"errors"
	"regexp"
	"time"
)

// WeekStartDay set week start day, default is sunday
var WeekStartDay = time.Sunday

// TimeFormats default time formats will be parsed as
var TimeFormats = []string{
	"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2",
	"15:4:5", "15:4", "15",
	"15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05Z0700", "2006-01-02T15:04:05Z07",
	"2006.1.2", "2006.1.2 15:04:05", "2006.01.02", "2006.01.02 15:04:05", "2006.01.02 15:04:05.999999999",
	"1/2/2006", "1/2/2006 15:4:5", "2006/01/02", "20060102", "2006/01/02 15:04:05",
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
	time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

// Config configuration for now package
type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

// DefaultConfig default config
var DefaultConfig *Config

// New initialize Now based on configuration
func (config *Config) With(t time.Time) *Now {
	return &Now{Time: t, Config: config}
}

// Parse parse string to time based on configuration
func (config *Config) Parse(strs ...string) (time.Time, error) {
	if config.TimeLocation == nil {
		return config.With(time.Now()).Parse(strs...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).Parse(strs...)
	}
}

// MustParse must parse string to time or will panic
func (config *Config) MustParse(strs ...string) time.Time {
	if config.TimeLocation == nil {
		return config.With(time.Now()).MustParse(strs...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).MustParse(strs...)
	}
}

// Now now struct
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
			TimeFormats:  TimeFormats,
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

func Monday(strs ...string) time.Time {
	return With(time.Now()).Monday(strs...)
}

// Sunday sunday
func Sunday(strs ...string) time.Time {
	return With(time.Now()).Sunday(strs...)
}

// EndOfSunday end of sunday
func EndOfSunday() time.Time {
	return With(time.Now()).EndOfSunday()
}

// Quarter returns the yearly quarter
func Quarter() uint {
	return With(time.Now()).Quarter()
}

// Parse parse string to time
func Parse(strs ...string) (time.Time, error) {
	return With(time.Now()).Parse(strs...)
}

// ParseInLocation parse string to time in location
func ParseInLocation(loc *time.Location, strs ...string) (time.Time, error) {
	return With(time.Now().In(loc)).Parse(strs...)
}

// MustParse must parse string to time or will panic
func MustParse(strs ...string) time.Time {
	return With(time.Now()).MustParse(strs...)
}

// MustParseInLocation must parse string to time in location or will panic
func MustParseInLocation(loc *time.Location, strs ...string) time.Time {
	return With(time.Now().In(loc)).MustParse(strs...)
}

// Between check now between the begin, end time or not
func Between(time1, time2 string) bool {
	return With(time.Now()).Between(time1, time2)
}

// BeginningOfMinute beginning of minute
func (now *Now) BeginningOfMinute() time.Time {
	return now.Truncate(time.Minute)
}

// BeginningOfHour beginning of hour
func (now *Now) BeginningOfHour() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, now.Time.Hour(), 0, 0, 0, now.Time.Location())
}

// BeginningOfDay beginning of day
func (now *Now) BeginningOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, now.Time.Location())
}

// BeginningOfWeek beginning of week
func (now *Now) BeginningOfWeek() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())

	if now.WeekStartDay != time.Sunday {
		weekStartDayInt := int(now.WeekStartDay)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}
	return t.AddDate(0, 0, -weekday)
}

// BeginningOfMonth beginning of month
func (now *Now) BeginningOfMonth() time.Time {
	y, m, _ := now.Date()
	return time.Date(y, m, 1, 0, 0, 0, 0, now.Location())
}

// BeginningOfQuarter beginning of quarter
func (now *Now) BeginningOfQuarter() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 3
	return month.AddDate(0, -offset, 0)
}

// BeginningOfHalf beginning of half year
func (now *Now) BeginningOfHalf() time.Time {
	month := now.BeginningOfMonth()
	offset := (int(month.Month()) - 1) % 6
	return month.AddDate(0, -offset, 0)
}

// BeginningOfYear BeginningOfYear beginning of year
func (now *Now) BeginningOfYear() time.Time {
	y, _, _ := now.Date()
	return time.Date(y, time.January, 1, 0, 0, 0, 0, now.Location())
}

// EndOfMinute end of minute
func (now *Now) EndOfMinute() time.Time {
	return now.BeginningOfMinute().Add(time.Minute - time.Nanosecond)
}

// EndOfHour end of hour
func (now *Now) EndOfHour() time.Time {
	return now.BeginningOfHour().Add(time.Hour - time.Nanosecond)
}

// EndOfDay end of day
func (now *Now) EndOfDay() time.Time {
	y, m, d := now.Date()
	return time.Date(y, m, d, 23, 59, 59, int(time.Second-time.Nanosecond), now.Location())
}

// EndOfWeek end of week
func (now *Now) EndOfWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// EndOfMonth end of month
func (now *Now) EndOfMonth() time.Time {
	return now.BeginningOfMonth().AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// EndOfQuarter end of quarter
func (now *Now) EndOfQuarter() time.Time {
	return now.BeginningOfQuarter().AddDate(0, 3, 0).Add(-time.Nanosecond)
}

// EndOfHalf end of half year
func (now *Now) EndOfHalf() time.Time {
	return now.BeginningOfHalf().AddDate(0, 6, 0).Add(-time.Nanosecond)
}

// EndOfYear end of year
func (now *Now) EndOfYear() time.Time {
	return now.BeginningOfYear().AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// Monday monday
/*
func (now *Now) Monday() time.Time {
	t := now.BeginningOfDay()
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return t.AddDate(0, 0, -weekday+1)
}
*/

func (now *Now) Monday(strs ...string) time.Time {
	var parseTime time.Time
	var err error
	if len(strs) > 0 {
		parseTime, err = now.Parse(strs...)
		if err != nil {
			panic(err)
		}
	} else {
		parseTime = now.BeginningOfDay()
	}
	weekday := int(parseTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return parseTime.AddDate(0, 0, -weekday+1)
}

func (now *Now) Sunday(strs ...string) time.Time {
	var parseTime time.Time
	var err error
	if len(strs) > 0 {
		parseTime, err = now.Parse(strs...)
		if err != nil {
			panic(err)
		}
	} else {
		parseTime = now.BeginningOfDay()
	}
	weekday := int(parseTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return parseTime.AddDate(0, 0, (7 - weekday))
}

// EndOfSunday end of sunday
func (now *Now) EndOfSunday() time.Time {
	return New(now.Sunday()).EndOfDay()
}

// Quarter returns the yearly quarter
func (now *Now) Quarter() uint {
	return (uint(now.Month())-1)/3 + 1
}

func (now *Now) parseWithFormat(str string, location *time.Location) (t time.Time, err error) {
	for _, format := range now.TimeFormats {
		t, err = time.ParseInLocation(format, str, location)

		if err == nil {
			return
		}
	}
	err = errors.New("Can't parse string as time: " + str)
	return
}

var hasTimeRegexp = regexp.MustCompile(`(\s+|^\s*|T)\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))(\s*$|[Z+-])`) // match 15:04:05, 15:04:05.000, 15:04:05.000000 15, 2017-01-01 15:04, 2021-07-20T00:59:10Z, 2021-07-20T00:59:10+08:00, 2021-07-20T00:00:10-07:00 etc
var onlyTimeRegexp = regexp.MustCompile(`^\s*\d{1,2}((:\d{1,2})*|((:\d{1,2}){2}\.(\d{3}|\d{6}|\d{9})))\s*$`)                // match 15:04:05, 15, 15:04:05.000, 15:04:05.000000, etc
func formatTimeToList(t time.Time) []int {
	hour, min, sec := t.Clock()
	year, month, day := t.Date()
	return []int{t.Nanosecond(), sec, min, hour, day, int(month), year}
}

// Parse parse string to time
func (now *Now) Parse(strs ...string) (t time.Time, err error) {
	var (
		setCurrentTime  bool
		parseTime       []int
		currentLocation = now.Location()
		onlyTimeInStr   = true
		currentTime     = formatTimeToList(now.Time)
	)

	for _, str := range strs {
		hasTimeInStr := hasTimeRegexp.MatchString(str) // match 15:04:05, 15
		onlyTimeInStr = hasTimeInStr && onlyTimeInStr && onlyTimeRegexp.MatchString(str)
		if t, err = now.parseWithFormat(str, currentLocation); err == nil {
			location := t.Location()
			parseTime = formatTimeToList(t)

			for i, v := range parseTime {
				// Don't reset hour, minute, second if current time str including time
				if hasTimeInStr && i <= 3 {
					continue
				}

				// If value is zero, replace it with current time
				if v == 0 {
					if setCurrentTime {
						parseTime[i] = currentTime[i]
					}
				} else {
					setCurrentTime = true
				}

				// if current time only includes time, should change day, month to current time
				if onlyTimeInStr {
					if i == 4 || i == 5 {
						parseTime[i] = currentTime[i]
						continue
					}
				}
			}

			t = time.Date(parseTime[6], time.Month(parseTime[5]), parseTime[4], parseTime[3], parseTime[2], parseTime[1], parseTime[0], location)
			currentTime = formatTimeToList(t)
		}
	}
	return
}

// MustParse must parse string to time or it will panic
func (now *Now) MustParse(strs ...string) (t time.Time) {
	t, err := now.Parse(strs...)
	if err != nil {
		panic(err)
	}
	return t
}

// Between check time between the begin, end time or not
func (now *Now) Between(begin, end string) bool {
	beginTime := now.MustParse(begin)
	endTime := now.MustParse(end)
	return now.After(beginTime) && now.Before(endTime)
}
