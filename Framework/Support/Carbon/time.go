package Carbon

import (
	"strconv"
	"strings"
	"time"
)

// Time - Similar function of time() in PHP.
//
// Original : https://www.php.net/manual/en/function.time.php
//
// Returns the current time measured in the number of seconds since the Unix Epoch (January 1 1970 00:00:00 GMT).
func Time() int64 {
	return time.Now().Unix()
}

//Unix 1654041600
func Unix(date time.Time) int64 {
	return date.Unix()
}
func UnixMilli(date time.Time) int64 {
	return date.UnixMilli()
}
func UnixMicro(date time.Time) int64 {
	return date.UnixMicro()
}
func UnixNano(date time.Time) int64 {
	return date.UnixNano()
}
func FromTimestamp(SecondTimestamp int64) time.Time {
	return time.Unix(SecondTimestamp, 0)
}
func FromMicrosecondTimestamp(MicrosecondTimestamp int64) time.Time {
	return time.UnixMicro(MicrosecondTimestamp)
}

func FromString(timeString string) time.Time {
	//timeString := "1657861095"
	timestamp, err := strconv.ParseInt(timeString, 10, 64)
	if err != nil {
		panic(err)
	}
	return time.Unix(timestamp, 0)
}

// internal utility methods
func webTime(t time.Time) string {
	ftime := t.Format(time.RFC1123)
	if strings.HasSuffix(ftime, "UTC") {
		ftime = ftime[0:len(ftime)-3] + "GMT"
	}
	return ftime
}

// Date date()
// Date("02/01/2006 15:04:05 PM", 1524799394)
// Note: the behavior is inconsistent with php's date function
func Date(format string, timestamp int64) string {
	return time.Unix(timestamp, 0).Format(format)
}
