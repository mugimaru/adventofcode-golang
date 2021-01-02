package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("12", SolveDay12)
}

const (
	E = 0
	S = 1
	W = 2
	N = 3
)

type SheepInstruction struct {
	Cmd   rune
	Value int
}

func SolveDay12(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	instructions := make([]SheepInstruction, len(lines))

	for i, line := range lines {
		rline := []rune(line)
		value, _ := strconv.Atoi(string(rline[1:]))

		switch rline[0] {
		case 'S', 'W':
			value = -value
		}

		instructions[i] = SheepInstruction{rline[0], value}
	}

	// direction := 'E'
	dir := 0
	posE := 0
	posN := 0

	for _, inst := range instructions {
		switch inst.Cmd {
		case 'N', 'S':
			posN += inst.Value
		case 'W', 'E':
			posE += inst.Value
		case 'R':
			dir = (dir + inst.Value/90) % 4
		case 'L':
			dir = (dir - inst.Value/90 + 4) % 4
		case 'F':
			switch dir {
			case E:
				posE += inst.Value
			case W:
				posE -= inst.Value
			case N:
				posN += inst.Value
			case S:
				posN -= inst.Value
			}
		}
	}

	dist1 := abs(posE) + abs(posN)

	// part 2
	posE = 0
	posN = 0
	wpE := 10
	wpN := 1

	for _, inst := range instructions {
		switch inst.Cmd {
		case 'F':
			posE += wpE * inst.Value
			posN += wpN * inst.Value
		case 'L':
			for i := 0; i < inst.Value/90; i++ {
				wpE, wpN = -wpN, wpE
			}
		case 'R':
			for i := 0; i < inst.Value/90; i++ {
				wpE, wpN = wpN, -wpE
			}
		case 'N', 'S':
			wpN += inst.Value
		case 'W', 'E':
			wpE += inst.Value
		}
	}

	dist2 := abs(posE) + abs(posN)

	return dist1, dist2
}

func abs(v int) int {
	if v < 0 {
		return -v
	}

	return v
}
