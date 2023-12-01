package aoc2023

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	input := map[string]int{
		"two1nine":         29,
		"eightwothree":     83,
		"abcone2threexyz":  13,
		"xtwone3four":      24,
		"4nineeightseven2": 42,
		"zoneight234":      14,
		"7pqrstsixteen":    76,
	}

	for line, expectedResult := range input {
		result := findNumberDay01(line, digitsRegexpPart2, digitsRegexpRevPart2)
		if result != expectedResult {
			t.Errorf("for line %s found %d instead of %d", line, result, expectedResult)
		}
	}
}
