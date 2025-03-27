package utils

import "time"

func ParseDate(date string) (time.Time, error) {
	return time.Parse("Mon, 02 Jan 2006 15:04:05 GMT", date)
}
