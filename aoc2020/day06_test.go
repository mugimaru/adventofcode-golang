package aoc2020

import (
	"testing"
)

func TestDay6(t *testing.T) {
	var input = `abc

a
b
c

ab
ac

a
a
a
a

b`

	p1, p2 := SolveDay06(input)

	if p1 != 11 || p2 != 6 {
		t.Errorf("expected p1=11 p2=6 got p1=%v p2=%v", p1, p2)
	}
}
