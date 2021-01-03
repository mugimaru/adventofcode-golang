package aoc2020

import (
	"testing"
)

func TestDay16(t *testing.T) {
	p1, _ := SolveDay16(`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`)

	if p1 != 71 {
		t.Errorf("expected p1=51  got p1=%v", p1)
	}
}
