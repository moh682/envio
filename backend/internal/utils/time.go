package utils

import (
	"log"
	"time"
)

func ParseFromUTC(s string) (time.Time, error) {
	log.Println("paring time -------- -- -- -- - -- - - ", s)
	return time.Parse("2006-01-02T15:04:05Z", s)
}

func IsBetweenTimes(t, start, end time.Time) bool {
	return t.After(start) && t.Before(end)
}

func IsYearAndMonthEqual(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month()
}
