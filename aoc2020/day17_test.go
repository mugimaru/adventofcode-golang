package aoc2020

import (
	"testing"
)

func TestDay17(t *testing.T) {
	p1, p2 := SolveDay17(`.#.
..#
###`)

	if p1 != 112 || p2 != 848 {
		t.Errorf("expected p1=112 p2=848  got p1=%v p2=%v", p1, p2)
	}
}
