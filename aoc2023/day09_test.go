package aoc2023

import (
	"testing"
)

func TestDay09Part1(t *testing.T) {
	p1, p2 := solveDay09(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`)

	if p1 != 114 {
		t.Errorf("p1 expected to eq 114 got %d", p1)
	}

	if p2 != 2 {
		t.Errorf("p2 expected to eq 2 got %d", p2)
	}
}
