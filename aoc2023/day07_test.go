package aoc2023

import (
	"testing"
)

var day07TestInput string = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestDay07(t *testing.T) {
	p1, p2 := SolveDay07(day07TestInput)

	if p1 != 6440 {
		t.Errorf("p1 expected to eq 6440 got %d", p1)
	}

	if p2 != 5905 {
		t.Errorf("p2 expected to eq 5905 got %d", p1)
	}
}
