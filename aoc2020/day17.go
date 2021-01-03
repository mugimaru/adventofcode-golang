package aoc2020

import (
	"strings"
)

func init() {
	registerFun("17", SolveDay17)
}

type point3d struct {
	x int
	y int
	z int
}

type point4d struct {
	x int
	y int
	z int
	w int
}

type cube struct {
	isActive   bool
	neighbours int
}

func countNeighbours4d(cubes map[point4d]cube) {
	for p, cube := range cubes {
		if !cube.isActive {
			continue
		}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						n := point4d{p.x + i, p.y + j, p.z + k, p.w + l}
						if n != p {
							nCube := cubes[n]
							nCube.neighbours++
							cubes[n] = nCube
						}
					}
				}
			}
		}
	}
}

func runCycle4d(cubes map[point4d]cube) {
	for key, c := range cubes {
		if c.isActive && (c.neighbours < 2 || c.neighbours > 3) || !c.isActive && c.neighbours != 3 {
			delete(cubes, key)
		} else {
			cubes[key] = cube{true, 0}
		}
	}
}

func countNeighbours(cubes map[point3d]cube) {
	for p, cube := range cubes {
		if !cube.isActive {
			continue
		}

		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					n := point3d{p.x + i, p.y + j, p.z + k}
					if n != p {
						nCube := cubes[n]
						nCube.neighbours++
						cubes[n] = nCube
					}
				}
			}
		}
	}
}

func runCycle(cubes map[point3d]cube) {
	for key, c := range cubes {
		if c.isActive && (c.neighbours < 2 || c.neighbours > 3) || !c.isActive && c.neighbours != 3 {
			delete(cubes, key)
		} else {
			cubes[key] = cube{true, 0}
		}
	}
}

func SolveDay17(input string) (interface{}, interface{}) {
	cubes := make(map[point3d]cube)
	cubes4d := make(map[point4d]cube)

	for y, line := range strings.Split(input, "\n") {
		for x, state := range []rune(line) {
			if state == '#' {
				cubes4d[point4d{x, y, 0, 0}] = cube{true, 0}
				cubes[point3d{x, y, 0}] = cube{true, 0}
			}
		}
	}

	for cycle := 0; cycle < 6; cycle++ {
		countNeighbours(cubes)
		runCycle(cubes)
		// TODO: refactor
		countNeighbours4d(cubes4d)
		runCycle4d(cubes4d)
	}

	return len(cubes), len(cubes4d)
}
