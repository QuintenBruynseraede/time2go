package utils

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/timerange"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatAsJSList[T any](list []T, quotes bool) template.JS {
	stringList := []string{}
	for _, x := range list {
		if quotes {
			stringList = append(stringList, fmt.Sprintf("\"%v\"", x))
		} else {
			stringList = append(stringList, fmt.Sprintf("%v", x))
		}
	}

	return template.JS("[" + strings.Join(stringList, ",") + "]")
}

func GenerateTimeRangeLabels(timeRange timerange.TimeRange) []template.JS {
	i := timeRange.Start
	labels := []template.JS{}

	for i.Before(timeRange.End) {
		labels = append(labels, template.JS(i.Format("Mon 15:04")))
		i = i.Add(time.Hour)
	}

	return labels
}

// Title capitalizes the first letter of each word
func Title(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}
