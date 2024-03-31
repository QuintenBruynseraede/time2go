package weather

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHighestConsecutiveScores(t *testing.T) {
	scores := []int{1, 2, 4}
	length := 3
	assert.Equal(t, []int{0, 1, 2}, findHighestConsecutiveScores(scores, length))
}

func TestGetScoreTotal(t *testing.T) {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{2, 3, 4, 5, 6}
	l3 := []int{3, 4, 5, 6, 7}
	out := getScoreTotal(l1, l2, l3)

	assert.Equal(t, []int{6, 9, 12, 15, 18}, out)
}
