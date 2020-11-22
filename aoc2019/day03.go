package aoc2019

import (
	"strconv"
	"strings"
)

type step struct {
	dir string
	len int
}

type point struct {
	x int
	y int
}

func init() {
	registerFun("03", SolveDay03)
}

func SolveDay03(input string) (interface{}, interface{}) {
	firstPath, secondPath := parseInput(input)

	var m = make(map[point]int)

	walk(firstPath, m, false)
	intersections, totalSteps := walk(secondPath, m, true)

	minSteps := totalSteps[0]
	minDistance := calculateDistance(intersections[0])
	for i := 1; i < len(intersections); i++ {
		distance := calculateDistance(intersections[i])
		if minDistance > distance {
			minDistance = distance
		}

		if minSteps > totalSteps[i] {
			minSteps = totalSteps[i]
		}
	}

	return minDistance, minSteps
}

func calculateDistance(p point) int {
	return abs(p.x) + abs(p.y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func walk(plan []step, walked map[point]int, freeze bool) ([]point, []int) {
	steps := 0
	currentPoint := point{x: 0, y: 0}
	var intersections []point
	var totalSteps []int

	for _, step := range plan {
		for i := 0; i < step.len; i++ {
			steps++
			currentPoint = nextPoint(currentPoint, step.dir)

			if walked[currentPoint] > 0 {
				intersections = append(intersections, currentPoint)
				totalSteps = append(totalSteps, walked[currentPoint]+steps)
			} else {
				if !freeze {
					walked[currentPoint] = steps
				}
			}
		}
	}

	return intersections, totalSteps
}

func nextPoint(currentPoint point, dir string) point {
	newPoint := point{x: currentPoint.x, y: currentPoint.y}

	switch dir {
	case "L":
		newPoint.x--
	case "R":
		newPoint.x++
	case "U":
		newPoint.y++
	case "D":
		newPoint.y--
	}

	return newPoint
}

func parseInput(rawInput string) ([]step, []step) {
	pathsStr := strings.Split(rawInput, "\n")
	return parsePath(pathsStr[0]), parsePath(pathsStr[1])
}

func parsePath(pathStr string) []step {
	path := strings.Split(pathStr, ",")
	parsed := make([]step, len(path))

	for i, st := range path {
		len, _ := strconv.Atoi(st[1:])
		parsed[i] = step{dir: st[0:1], len: len}
	}

	return parsed
}
