package aoc2023

import (
	"testing"
)

func TestDay08Part1(t *testing.T) {
	nodes, directions := parseDay08Input(`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`)

	res := stepsToTarget("AAA", nodes, directions, isEndOfPath)
	if res != 6 {
		t.Errorf("p1 expected to eq 6 got %d", res)
	}
}

func TestDay08Part2(t *testing.T) {
	nodes, directions := parseDay08Input(`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`)
	res := solveDay08Part2(nodes, directions)
	if res != 6 {
		t.Errorf("p2 expected to eq 6 got %d", res)
	}
}
