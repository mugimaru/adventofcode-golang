package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registerFun("01", SolveDay01)
}

func SolveDay01(input string) (interface{}, interface{}) {
	expensesReport := parseInput(input)
	fmt.Println(expensesReport)
	return solvePart1(expensesReport), solvePart2(expensesReport)
}

func solvePart1(report []int) int {
	for _, a := range report {
		for _, b := range report {
			if a+b == 2020 {
				return a * b
			}
		}
	}

	panic("solution not found")
}

func solvePart2(report []int) int {
	for _, a := range report {
		for _, b := range report {
			for _, c := range report {
				if a+b+c == 2020 {
					return a * b * c
				}
			}

		}
	}

	panic("solution not found")
}

func parseInput(input string) []int {
	lines := strings.Split(input, "\n")
	result := make([]int, len(lines))

	for i := 0; i < len(lines); i++ {
		value, err := strconv.Atoi(lines[i])
		if err != nil {
			panic(err)
		}
		result[i] = value
	}

	return result
}
