package aoc2023

import (
	"testing"
)

var day05TestInput string = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestDay05(t *testing.T) {
	p1, p2 := SolveDay05(day05TestInput)

	if p1 != 35 {
		t.Errorf("p1 expected to eq 35 got %d", p1)
	}

	if p2 != 46 {
		t.Errorf("p2 expected to eq 46 got %d", p2)
	}
}
