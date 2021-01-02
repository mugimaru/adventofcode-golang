package aoc2020

import (
	"testing"
)

func TestDay11(t *testing.T) {
	var input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

	p1, p2 := SolveDay11(input)

	if p1 != 37 && p2 != 26 {
		t.Errorf("expected p1=37 p2=26 got p1=%v, p2=%v", p1, p2)
	}
}
