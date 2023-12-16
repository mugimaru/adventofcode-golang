package aoc2023

import (
	"github.com/mugimaru73/adventofcode-golang/utils"
	"strings"
)

func init() {
	registerFun("09", solveDay09)
}

func solveDay09(input string) (interface{}, interface{}) {
	report := parseDay09Input(input)

	var p1Sum, p2Sum int
	for _, line := range report {
		p1Inc, p2Inc := nextValueForReportLine(line)
		p1Sum += p1Inc
		p2Sum += p2Inc
	}
	return p1Sum, p2Sum
}

func nextValueForReportLine(input []int) (int, int) {
	lastVals := []int{input[len(input)-1]}
	firstVals := []int{input[0]}
	currentLine := input
	for {
		diff := make([]int, 0)
		allZeroes := true
		for i := 1; i < len(currentLine); i++ {
			d := currentLine[i] - currentLine[i-1]
			if d != 0 {
				allZeroes = false
			}
			diff = append(diff, d)
		}
		firstVals = append(firstVals, diff[0])
		lastVals = append(lastVals, diff[len(diff)-1])
		currentLine = diff

		if allZeroes {
			break
		}
	}

	p2 := firstVals[len(firstVals)-1]
	for i := len(firstVals) - 2; i >= 0; i-- {
		p2 = firstVals[i] - p2
	}

	return utils.SumIntSlice(lastVals), p2
}

func parseDay09Input(input string) (res [][]int) {
	for _, line := range strings.Split(input, "\n") {
		res = append(res, utils.ParseNumsRow(line))
	}
	return res
}
