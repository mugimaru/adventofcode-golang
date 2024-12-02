package aoc2024

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func init() {
	registerFun("01", SolveDay01)
}

func SolveDay01(input string) (interface{}, interface{}) {
	a, b := parseInputDay01(input)

	slices.Sort(a)
	slices.Sort(b)
	var diff float64
	for i := 0; i < len(a); i++ {
		diff += math.Abs(float64(a[i] - b[i]))
	}

	freq := make(map[int]int)
	for _, item := range a {
		freq[item] = 0
	}

	for _, item := range b {
		if _, ok := freq[item]; ok {
			freq[item]++
		}
	}

	var simScore int
	for _, item := range a {
		simScore += item * freq[item]
	}

	return diff, simScore
}

func parseInputDay01(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")
	a := make([]int, len(lines))
	b := make([]int, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line)
		aVal, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		a[i] = aVal

		bVal, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		b[i] = bVal
	}

	return a, b
}
