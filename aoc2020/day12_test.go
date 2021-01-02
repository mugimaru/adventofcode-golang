package aoc2020

import (
	"testing"
)

func TestDay12(t *testing.T) {
	var input = `F10
N3
F7
R90
F11`

	p1, p2 := SolveDay12(input)

	if p1 != 25 && p2 != 286 {
		t.Errorf("expected p1=25 p2=286 got p1=%v, p2=%v", p1, p2)
	}
}
