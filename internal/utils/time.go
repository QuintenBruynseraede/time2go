package utils

import (
	"errors"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/timerange"
)

func DateRangeInputToTimeRange(rangeString string) (timerange.TimeRange, error) {
	// rangeString can be "today", "tomorrow" or "this_week"
	today := time.Now()
	if rangeString == "today" {
		return timerange.TimeRange{
			Start: startOfDayPlusN(today, 0),
			End:   startOfDayPlusN(today, 1),
		}, nil
	} else if rangeString == "tomorrow" {
		return timerange.TimeRange{
			Start: startOfDayPlusN(today, 1),
			End:   startOfDayPlusN(today, 2),
		}, nil
	} else if rangeString == "this_week" {
		return timerange.TimeRange{
			Start: startOfDayPlusN(today, 0),
			End:   startOfDayPlusN(today, 7),
		}, nil
	}

	return timerange.TimeRange{}, errors.New("invalid date range")
}

func Truncate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func startOfDayPlusN(day time.Time, n int) time.Time {
	return Truncate(day).Add(time.Hour * time.Duration(n*24))
}
