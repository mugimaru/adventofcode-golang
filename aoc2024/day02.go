package aoc2024

import (
	"cmp"
	"github.com/mugimaru73/adventofcode-golang/utils"
	"math"
	"strconv"
	"strings"
)

func init() {
	registerFun("02", SolveDay02)
}

func SolveDay02(input string) (interface{}, interface{}) {
	reports := parseInputDay02(input)
	var numSafe, numSafePart2 int
	for _, report := range reports {
		if isReportSafe(report) {
			numSafe++
			numSafePart2++
			continue
		}
		if isReportSafeP2(report) {
			numSafePart2++
		}

	}
	return numSafe, numSafePart2
}

func isReportSafeP2(report []int) bool {
	for i := 0; i < len(report); i++ {
		newReport := utils.RemoveElementByIndex(report, i)
		if isReportSafe(newReport) {
			return true
		}
	}

	return false
}

func isReportSafe(report []int) bool {
	var dir int
	for i := 1; i < len(report); i++ {
		diff := math.Abs(float64(report[i] - report[i-1]))
		if diff < 1 || diff > 3 {
			return false
		}

		cmpRes := cmp.Compare(report[i], report[i-1])
		if dir == 0 {
			dir = cmpRes
		}
		if cmpRes != 0 && cmpRes != dir {
			return false
		}

	}

	return true
}

func parseInputDay02(input string) [][]int {
	reports := make([][]int, 0)
	for _, line := range strings.Split(input, "\n") {
		itemsStr := strings.Fields(line)
		report := make([]int, len(itemsStr))
		for i, s := range itemsStr {
			item, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			report[i] = item
		}

		reports = append(reports, report)
	}
	return reports
}
