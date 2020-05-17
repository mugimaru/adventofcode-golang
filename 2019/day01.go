package main

import (
	"math"
	"strconv"
	"strings"

	"github.com/mugimaru73/adventofcode/executor"
)

func run(input string) (interface{}, interface{}) {
	var Part1 = 0.0
	var Part2 = 0.0

	for _, v := range strings.Split(input, "\n") {
		distance, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}

		fuelReq := fuelRequired(distance)
		Part1 += fuelReq
		Part2 += part2FuelRequired(fuelReq, fuelReq)
	}

	return int(Part1), int(Part2)
}

func fuelRequired(distance float64) float64 {
	return math.Trunc(distance/3) - 2.0
}

func part2FuelRequired(acc float64, distance float64) float64 {
	var req = fuelRequired(distance)
	if req < 0 {
		req = 0
	}

	if distance-req > 0 {
		acc = part2FuelRequired(acc+req, req)
	}

	return acc
}

func main() {
	executor.Run(executor.ReadInput("2019/day01.input.txt"), run)
}
