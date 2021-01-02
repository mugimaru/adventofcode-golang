package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("08", SolveDay08)
}

type Cmd struct {
	cmd   string
	value int
}

func SolveDay08(input string) (interface{}, interface{}) {
	program := parseProgram(input)
	p1, _ := exec(program)
	return p1, fixProgram(program)
}

func fixProgram(program []Cmd) int {
	for i, cmd := range program {
		if cmd.cmd != "acc" {
			if value, hasInfLoop := exec(changeCommand(program, i)); !hasInfLoop {
				return value
			}
		}
	}

	panic(":(")
}

func parseProgram(input string) []Cmd {
	lines := strings.Split(input, "\n")
	program := make([]Cmd, len(lines))

	for i, line := range lines {
		cmdr := []rune(line)

		value, err := strconv.Atoi(string(cmdr[4:]))
		if err != nil {
			panic(err)
		}

		program[i] = Cmd{cmd: string(cmdr[0:3]), value: value}
	}

	return program
}

func changeCommand(program []Cmd, index int) []Cmd {
	copy := make([]Cmd, len(program))

	for i := 0; i < len(program); i++ {
		copy[i] = program[i]
	}

	switch copy[index].cmd {
	case "nop":
		copy[index].cmd = "jmp"
	case "jmp":
		copy[index].cmd = "nop"
	default:
		panic(":(")
	}

	return copy
}

func exec(program []Cmd) (int, bool) {
	acc := 0
	i := 0
	executed := make(map[int]bool)

	len := len(program)

	for {
		if executed[i] {
			return acc, true
		}

		if i == len {
			return acc, false
		}

		executed[i] = true

		switch program[i].cmd {
		case "nop":
			i++
		case "acc":
			acc += program[i].value
			i++
		case "jmp":
			i += program[i].value
		}
	}
}
