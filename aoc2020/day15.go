package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("15", SolveDay15)
}

func SolveDay15(input string) (interface{}, interface{}) {
	return day15PlayGame(input, 2020), day15PlayGame(input, 30000000)
}

func day15PlayGame(input string, lastTurn int) int {
	var lastSpokenNum int

	history := make(map[int]int)

	for turn, v := range strings.Split(input, ",") {
		lastSpokenNum, _ = strconv.Atoi(v)
		history[lastSpokenNum] = turn
	}

	for turn := len(history) + 1; turn < lastTurn; turn++ {
		lastSpokenTurn, spokenBefore := history[lastSpokenNum]

		if spokenBefore {
			history[lastSpokenNum] = turn - 1
			lastSpokenNum = turn - 1 - lastSpokenTurn
		} else {
			history[lastSpokenNum] = turn - 1
			lastSpokenNum = 0
		}
	}

	return lastSpokenNum
}
