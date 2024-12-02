package aoc2024

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay01Part1(t *testing.T) {
	input := `3   4
4   3
2   5
1   3
3   9
3   3`

	p1, p2 := SolveDay01(input)
	assert.Equal(t, float64(11), p1)
	assert.Equal(t, 31, p2)
}
