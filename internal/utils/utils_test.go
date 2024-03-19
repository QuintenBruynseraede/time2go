package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatIntList(t *testing.T) {
	list := []interface{}{1, 2, 3}

	assert.Equal(t, FormatAsJSList(list, false), "[1,2,3]")
	assert.Equal(t, FormatAsJSList(list, true), `["1"","2","3"]`)
}

func TestFormatStringList(t *testing.T) {
	list := []interface{}{"a", "b"}

	assert.Equal(t, FormatAsJSList(list, false), "[a,b]")
	assert.Equal(t, FormatAsJSList(list, true), `["a"","b"]`)
}
