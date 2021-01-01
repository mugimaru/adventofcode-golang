package aoc2020

import (
	"strings"
)

func init() {
	registerFun("03", SolveDay03)
}

func SolveDay03(input string) (interface{}, interface{}) {
	treesMap := buildTreesMap(input)

	resultP1 := treesOnRoute(treesMap, 3, 1)
	resultP2 := resultP1
	for _, v := range [4][2]int{{1, 1}, {5, 1}, {7, 1}, {1, 2}} {
		resultP2 = resultP2 * treesOnRoute(treesMap, v[0], v[1])
	}

	return resultP1, resultP2
}

func buildTreesMap(input string) [][]bool {
	lines := strings.Split(input, "\n")
	lineLen := len(lines[0])
	treesMap := make([][]bool, len(lines))

	for i, line := range lines {
		hasTree := make([]bool, lineLen)
		for j, v := range []rune(line) {
			if string(v) == "#" {
				hasTree[j] = true
			} else {
				hasTree[j] = false
			}
		}
		treesMap[i] = hasTree
	}

	return treesMap
}

func treesOnRoute(treesMap [][]bool, stepX int, stepY int) int {
	x := 0
	y := 0
	treesOnRoute := 0
	lineLen := len(treesMap[0])

	for y < len(treesMap)-1 {
		x += stepX
		y += stepY

		normalizedX := x
		if x >= lineLen {
			normalizedX = x % lineLen
		}

		if treesMap[y][normalizedX] {
			treesOnRoute++
		}
	}

	return treesOnRoute
}
