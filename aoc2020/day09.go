package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("09", SolveDay09)
}

func SolveDay09(input string) (interface{}, interface{}) {
	numbers := parseDay09Input(input)
	invalidNum := findFirstInvalidNumber(numbers, 25)
	return invalidNum, findEncryptionWeakness(numbers, invalidNum)
}

func findEncryptionWeakness(numbers []int, invalidNum int) int {
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		min := numbers[i]
		max := numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if numbers[j] > max {
				max = numbers[j]
			}

			if numbers[j] < min {
				min = numbers[j]
			}

			if sum == invalidNum {
				return max + min
			}

			if sum > invalidNum {
				break
			}
		}
	}

	panic("not found")
}

func findFirstInvalidNumber(numbers []int, preamble int) int {
	for i := preamble; i < len(numbers); i++ {
		if _, hasPair := hasPair(numbers[i-preamble:i], numbers[i]); !hasPair {
			return numbers[i]
		}
	}

	return 0
}

func hasPair(numbers []int, value int) ([2]int, bool) {
	for i, iv := range numbers {
		for j, jv := range numbers {
			if i == j {
				continue
			}

			if iv+jv == value {
				return [2]int{iv, jv}, true
			}
		}
	}

	return [2]int{0, 0}, false
}

func parseDay09Input(input string) []int {
	lines := strings.Split(input, "\n")
	numbers := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		num, _ := strconv.Atoi(lines[i])
		numbers[i] = num
	}

	return numbers
}
