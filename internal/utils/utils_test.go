package utils

import (
	"testing"
	"time"

	"github.com/QuintenBruynseraede/time2go/internal/timerange"
	"github.com/stretchr/testify/assert"
)

func TestFormatIntList(t *testing.T) {
	list := []interface{}{1, 2, 3}

	assert.Equal(t, FormatAsJSList(list, false), "[1,2,3]")
	assert.Equal(t, FormatAsJSList(list, true), `["1","2","3"]`)
}

func TestFormatStringList(t *testing.T) {
	list := []interface{}{"a", "b"}

	assert.Equal(t, FormatAsJSList(list, false), "[a,b]")
	assert.Equal(t, FormatAsJSList(list, true), `["a","b"]`)
}

func TestStartOfDayPlusN(t *testing.T) {
	day := time.Date(2024, 1, 1, 5, 0, 0, 0, time.UTC)
	startOfDay := Truncate(day)
	dayPlus1 := startOfDay.Add(24 * time.Hour)
	dayMinus1 := startOfDay.Add(-24 * time.Hour)

	assert.Equal(t, startOfDayPlusN(day, 0), startOfDay)
	assert.Equal(t, startOfDayPlusN(day, 1), dayPlus1)
	assert.Equal(t, startOfDayPlusN(day, -1), dayMinus1)
}

func TestGenerateTimeRangeLabels(t *testing.T) {
	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC) // Monday 1/1/2024
	end := start.Add(4 * time.Hour)

	labels := GenerateTimeRangeLabels(timerange.TimeRange{Start: start, End: end})

	assert.Equal(t, []string{"Mon 00:00", "Mon 01:00", "Mon 02:00", "Mon 03:00"}, labels)
}
