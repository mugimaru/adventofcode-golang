package aoc2020

import (
	"testing"
)

func TestDay13(t *testing.T) {
	p1, p2 := SolveDay13("939\n7,13,x,x,59,x,31,19")

	if p1 != 295 || p2 != 100000000385544 {
		t.Errorf("expected p1=295 p2=100000000385544 got p1=%v, p2=%v", p1, p2)
	}
}
