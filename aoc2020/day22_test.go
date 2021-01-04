package aoc2020

import (
	"testing"
)

func TestDay22(t *testing.T) {
	input := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

	p1, p2 := SolveDay22(input)
	if p1 != 306 || p2 != 291 {
		t.Errorf("expected p1=306 p2=291 p1=%v p2=`%v`", p1, p2)
	}
}
