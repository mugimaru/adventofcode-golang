package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/mugimaru73/adventofcode-golang/executor"
)

func run(input string) (interface{}, interface{}) {
	orbits := make(map[string][]string)

	for _, v := range strings.Split(input, "\n") {
		objects := strings.Split(v, ")")
		orbits[objects[1]] = append(orbits[objects[1]], objects[0])
	}

	return countOrbitsSum(orbits), minOrbitalTransferLen(orbits["YOU"][0], orbits["SAN"][0], orbits)
}

func minOrbitalTransferLen(from string, to string, orbits map[string][]string) int {
	fromPaths := allPathsFrom(from, []string{}, orbits)
	toPaths := allPathsFrom(to, []string{}, orbits)

	distance := math.MaxInt32
	for _, path := range fromPaths {
		for i, fromPathPlanet := range path {
			for _, toPlanetPath := range toPaths {
				for j, toPathPlanet := range toPlanetPath {
					if fromPathPlanet == toPathPlanet {
						if newDistance := i + j; newDistance < distance {
							distance = newDistance
						}
					}
				}
			}
		}
	}

	return distance
}

func allPathsFrom(planet string, path []string, orbits map[string][]string) [][]string {
	var allPaths [][]string
	path = append(path, planet)

	if orbits[planet] == nil {
		return [][]string{path}
	}

	paths := make([][]string, len(orbits[planet]))
	for i := range paths {
		for _, newPath := range allPathsFrom(orbits[planet][i], path, orbits) {

			allPaths = append(allPaths, newPath)
		}
	}

	return allPaths
}

func countOrbitsSum(orbits map[string][]string) int {
	count := 0
	for from := range orbits {
		count += countOrbits(from, orbits)
	}

	return count
}

func countOrbits(from string, orbits map[string][]string) int {
	acc := 0
	if orbits[from] != nil {
		acc++
		for _, to := range orbits[from] {
			acc += countOrbits(to, orbits)
		}
	}

	return acc
}

func shortestPath(from string, to string, orbits map[string][]string, acc int) int {
	acc++
	fmt.Printf("from %v to %v:\n", from, to)

	minPath := math.MaxInt32

	for _, vertex := range orbits[from] {
		fmt.Printf("  test %v\n", vertex)
		if vertex == to {
			return acc
		}

		sp := shortestPath(vertex, to, orbits, acc)
		fmt.Printf("  sp from %v to %v is %v\n", vertex, to, sp)
		if sp < math.MaxInt32 {
			path := sp + acc
			if path < minPath {
				minPath = path
			}
		}
	}

	return minPath
}

func main() {
	executor.Run(executor.ReadInput("2019/day06.input.txt"), run)
}
