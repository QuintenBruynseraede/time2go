package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/timerange"
)

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

func GenerateTimeRangeLabels(timeRange timerange.TimeRange) []string {
	i := timeRange.Start
	labels := []string{}

	for i.Before(timeRange.End) {
		labels = append(labels, i.Format("Mon 15:04"))
		i = i.Add(time.Hour)
	}

	return labels
}
