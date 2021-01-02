package aoc2020

import (
	"testing"
)

func TestDay14(t *testing.T) {
	p1, p2 := SolveDay14(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`)

	if p1 != 51 || p2 != 208 {
		t.Errorf("expected p1=51 p2=208 got p1=%v, p2=%v", p1, p2)
	}
}
