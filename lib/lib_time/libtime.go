package libtime

import (
	"fmt"
	"strconv"
	"time"
)

func GetTime() string {
	now := time.Now()
	yyMMdd := fmt.Sprintf("%02d%02d%02d", now.Year()%100, int(now.Month()), now.Day())
	return fmt.Sprintf("%v", yyMMdd)
}

func ParseIntFormat(t time.Time) int64 {
	timeToInt := t.UnixNano() / int64(time.Second)

	if timeToInt < 0 {
		return 0
	}

	return timeToInt
}

func FormatTimeTypeDDT(date time.Time) string {
	return date.Format("Mon, 02 Jan 2006 15:04:05 GMT")
}

type Timestamp int64

// Get current time in milliseconds
func GetTimestamp() Timestamp {
	now := time.Now()
	return Timestamp(now.UnixNano() / int64(time.Millisecond))
}

// Convert timestamp to int64
func (t Timestamp) Int() int64 {
	return int64(t)
}

// Convert timestamp to string
func (t Timestamp) String() string {
	return strconv.FormatInt(t.Int(), 10)
}
