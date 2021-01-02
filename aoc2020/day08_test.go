package aoc2020

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	var input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

	p1, p2 := SolveDay08(input)

	if p1 != 5 {
		t.Errorf("expected p1=5 p2=8 got p1=%v, p2=%v", p1, p2)
	}
}
