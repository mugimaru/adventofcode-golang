package aoc2020

import (
	"testing"
)

const testInput = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

func TestDay3Solution(t *testing.T) {
	p1, p2 := SolveDay03(testInput)

	if p1 != 7 {
		t.Errorf("part1: Expected result is 7, but result=%v", p1)
	}
	if p2 != 336 {
		t.Errorf("part2: Expected result is 336, but result=%v", p2)
	}
}
