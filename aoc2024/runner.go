package aoc2024

import (
	"fmt"
	"time"
)

var funcMap = make(map[string]func(string) (interface{}, interface{}))

func registerFun(day string, fun func(string) (interface{}, interface{})) {
	funcMap[day] = fun
}

func Run(day string, input string) error {
	startedAt := time.Now()

	fun := funcMap[day]
	if fun == nil {
		return fmt.Errorf("aoc2024 doesn't implement a solution for day %v", day)
	}

	part1, part2 := fun(input)
	fmt.Printf("AdventOfCode 2024\nDay%v\n  Part1 = %v\n  Part2 = %v\nsolved in %s\n", day, part1, part2, time.Since(startedAt))

	return nil
}
