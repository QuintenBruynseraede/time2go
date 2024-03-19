package utils

import (
	"fmt"
	"strings"
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
