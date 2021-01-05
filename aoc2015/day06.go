package aoc2015

import (
	"fmt"
	"strings"
)

func init() {
	registerFun("06", solveDay06)
}

func solveDay06(input string) (interface{}, interface{}) {
	grid := make(map[point]bool)
	grid2 := make(map[point]int)

	for _, line := range strings.Split(input, "\n") {
		var fromX, fromY, toX, toY int

		fields := strings.Fields(strings.Replace(line, "turn ", "", 1))

		fmt.Sscanf(fields[1], "%d,%d", &fromX, &fromY)
		fmt.Sscanf(fields[3], "%d,%d", &toX, &toY)

		for y := fromY; y <= toY; y++ {
			for x := fromX; x <= toX; x++ {
				p := point{x, y}
				switch fields[0] {
				case "on":
					grid[p] = true
					grid2[p]++
				case "off":
					delete(grid, p)
					if grid2[p] > 0 {
						grid2[p]--
					}
				case "toggle":
					grid2[p] += 2
					if grid[p] {
						delete(grid, p)
					} else {
						grid[p] = true
					}
				}
			}
		}
	}

	brightness := 0
	for _, v := range grid2 {
		brightness += v
	}
	return len(grid), brightness
}
