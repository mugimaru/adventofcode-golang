package aoc2023

import (
	"testing"
)

var day03TestInput string = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDay03(t *testing.T) {
	p1, p2 := SolveDay03(day03TestInput)

	if p1 != 4361 {
		t.Errorf("p1 expected 4361 got %d", p1)
	}

	if p2 != 467835 {
		t.Errorf("p2 expected 467835 got %d", p2)
	}
}
