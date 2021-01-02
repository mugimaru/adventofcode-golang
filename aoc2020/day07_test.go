package aoc2020

import (
	"testing"
)

func TestDay7Part1(t *testing.T) {
	var input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	p1, p2 := SolveDay07(input)

	if p1 != 4 || p2 != 32 {
		t.Errorf("expected p1=4, p2=32 got p1=%v, p2=%v", p1, p2)
	}
}

func TestDay7Part2(t *testing.T) {
	var input = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

	_, p2 := SolveDay07(input)

	if p2 != 126 {
		t.Errorf("expected 126 got %v", p2)
	}
}
