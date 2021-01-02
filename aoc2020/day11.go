package aoc2020

import (
	"strings"
)

func init() {
	registerFun("11", SolveDay11)
}

func SolveDay11(input string) (interface{}, interface{}) {
	return simulateDay11(input, simulateRoundPart1), simulateDay11(input, simulateRoundPart2)
}

func countOccupiedSeatsInSight(sm [][]rune, x int, y int) int {
	count := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			cX := x
			cY := y
			for {
				cX += j
				cY += i

				if cX < 0 || cY < 0 || cX >= len(sm[0]) || cY >= len(sm) {
					break
				}

				if sm[cY][cX] == '#' {
					count++
					break
				}

				if sm[cY][cX] == 'L' {
					break
				}
			}

		}
	}

	return count
}

func countOccupiedAdjacentSeats(sm [][]rune, x int, y int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}

			cX := x + j
			cY := y + i

			if cX >= 0 && cY >= 0 && cX < len(sm[0]) && cY < len(sm) && sm[cY][cX] == '#' {
				count++
			}
		}
	}

	return count
}

func simulateRoundPart1(sm [][]rune) [][]rune {
	smCopy := make([][]rune, len(sm))

	for y := 0; y < len(sm); y++ {
		smCopy[y] = make([]rune, len(sm[y]))
		for x := 0; x < len(sm[y]); x++ {
			smCopy[y][x] = sm[y][x]
			if sm[y][x] == 'L' && countOccupiedAdjacentSeats(sm, x, y) == 0 {
				smCopy[y][x] = '#'
			}
			if sm[y][x] == '#' && countOccupiedAdjacentSeats(sm, x, y) >= 4 {
				smCopy[y][x] = 'L'
			}
		}
	}

	return smCopy
}

func simulateRoundPart2(sm [][]rune) [][]rune {
	smCopy := make([][]rune, len(sm))

	for y := 0; y < len(sm); y++ {
		smCopy[y] = make([]rune, len(sm[y]))
		for x := 0; x < len(sm[y]); x++ {
			smCopy[y][x] = sm[y][x]
			if sm[y][x] == 'L' && countOccupiedSeatsInSight(sm, x, y) == 0 {
				smCopy[y][x] = '#'
			}
			if sm[y][x] == '#' && countOccupiedSeatsInSight(sm, x, y) >= 5 {
				smCopy[y][x] = 'L'
			}
		}
	}

	return smCopy
}

func areSeatMapsEqual(sm1 [][]rune, sm2 [][]rune) bool {
	if len(sm1) != len(sm2) {
		return false
	}

	for i := 0; i < len(sm1); i++ {
		if len(sm1[i]) != len(sm2[i]) {
			return false
		}

		for j := 0; j < len(sm1[i]); j++ {
			if sm1[i][j] != sm2[i][j] {
				return false
			}
		}
	}

	return true
}

func simulateDay11(input string, fun func([][]rune) [][]rune) int {
	sm := parseSeatsMap(input)
	for {
		newSm := fun(sm)
		if areSeatMapsEqual(sm, newSm) {
			break
		}
		sm = newSm
	}

	occupiedSeats := 0
	for _, line := range sm {
		for _, seat := range line {
			if seat == '#' {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func parseSeatsMap(input string) [][]rune {
	lines := strings.Split(input, "\n")
	sm := make([][]rune, len(lines))

	for y, line := range lines {
		row := make([]rune, len(line))
		for x, value := range []rune(line) {
			row[x] = value
		}

		sm[y] = row
	}
	return sm
}
