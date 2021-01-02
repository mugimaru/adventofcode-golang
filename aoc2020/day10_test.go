package aoc2020

import (
	"testing"
)

func TestDay10(t *testing.T) {
	var input = `16
10
15
5
1
11
7
19
6
12
4`

	p1, p2 := SolveDay10(input)

	if p1 != 35 || p2 != 8 {
		t.Errorf("expected p1=35 p2=8 got p1=%v p2=%v", p1, p2)
	}
}
