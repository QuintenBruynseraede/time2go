package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/timerange"
)

func Repeat(value int, n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

func FormatAsJSList[T any](list []T, quotes bool) string {
	stringList := []string{}
	for _, x := range list {
		if quotes {
			stringList = append(stringList, fmt.Sprintf("\"%v\"", x))
		} else {
			stringList = append(stringList, fmt.Sprintf("%v", x))
		}
	}

	return "[" + strings.Join(stringList, ",") + "]"
}

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

func GenerateTimeRangeLabels(timeRange timerange.TimeRange) []string {
	i := timeRange.Start
	labels := []string{}

	for i.Before(timeRange.End) {
		labels = append(labels, i.Format("Mon 15:04"))
		i = i.Add(time.Hour)
	}

	return labels
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Between(min int, max int) []int {
	arr := []int{}
	for i := min; i < max; i++ {
		arr = append(arr, i)
	}
	return arr
}
