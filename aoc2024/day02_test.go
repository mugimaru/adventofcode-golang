package aoc2024

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDay02(t *testing.T) {
	input := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

	p1, p2 := SolveDay02(input)
	assert.Equal(t, 2, p1)
	assert.Equal(t, 4, p2)
}
