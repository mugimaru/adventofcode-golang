package aoc2023

import (
	"testing"
)

var day06TestInput string = `Time:      7  15   30
Distance:  9  40  200`

func TestDay06(t *testing.T) {
	p1, p2 := SolveDay06(day06TestInput)

	if p1 != 288 {
		t.Errorf("p1 expected to eq 288 got %d", p1)
	}

	if p2 != 71503 {
		t.Errorf("p2 expected to eq 71503 got %d", p2)
	}
}
