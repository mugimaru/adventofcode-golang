package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("13", SolveDay13)
}

func SolveDay13(input string) (interface{}, interface{}) {
	return solveDay13Part1(input), solveDay13Part2(input)
}

func solveDay13Part1(input string) int {
	inp := strings.Split(input, "\n")
	minDepartureTime, _ := strconv.Atoi(inp[0])

	var (
		buses    []int
		minDelay int
		busID    int
	)

	for _, v := range strings.Split(inp[1], ",") {
		if v != "x" {
			busID, _ := strconv.Atoi(v)
			buses = append(buses, busID)
		}
	}

	for i, bus := range buses {
		if bus-minDepartureTime%bus < minDelay || i == 0 {
			minDelay = bus - minDepartureTime%bus
			busID = bus
		}
	}

	return minDelay * busID
}

func solveDay13Part2(input string) int {
	inp := strings.Split(input, "\n")

	var (
		buses            []int
		minDepartureTime int = 100000000000000
		step             int = 1
	)

	for _, v := range strings.Split(inp[1], ",") {
		ts, _ := strconv.Atoi(v)
		buses = append(buses, ts)
	}

	for i, bus := range buses {
		if bus == 0 {
			continue
		}
		for (minDepartureTime+i)%bus != 0 {
			minDepartureTime += step
		}
		step *= bus
	}

	return minDepartureTime
}
