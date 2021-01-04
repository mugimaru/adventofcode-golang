package aoc2020

import (
	"testing"
)

func TestDay23(t *testing.T) {
	p1, p2 := SolveDay23("389125467")
	if p1 != "67384529" || p2 != 149245887792 {
		t.Errorf("expected p1=67384529 p2=149245887792 p1=%v p2=`%v`", p1, p2)
	}
}
