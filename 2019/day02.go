package main

import (
	"strconv"
	"strings"

	"github.com/mugimaru73/adventofcode-golang/executor"
)

var p2DesiredOutput = 19690720

func run(input string) (interface{}, interface{}) {
	program := readProgram(input)
	return calculateOutput(12, 10, program), solveP2(program)
}

func solveP2(program []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if calculateOutput(noun, verb, program) == p2DesiredOutput {
				return 100*noun + verb
			}
		}
	}

	panic("Solution not found")
}

func calculateOutput(noun int, verb int, program []int) int {
	p := make([]int, len(program))
	copy(p, program)
	p[1] = noun
	p[2] = verb
	return runProgram(0, p)[0]
}

func runProgram(i int, program []int) []int {
	switch opcode := program[i]; opcode {
	case 1:
		a1, a2, target := program[i+1], program[i+2], program[i+3]
		program[target] = program[a1] + program[a2]
		runProgram(i+4, program)
	case 2:
		a1, a2, target := program[i+1], program[i+2], program[i+3]
		program[target] = program[a1] * program[a2]
		runProgram(i+4, program)
	}

	return program
}

func dupProgram(program []int) []int {
	cp := make([]int, len(program))
	copy(cp, program)
	return cp
}

func readProgram(input string) []int {
	program := []int{}

	for _, v := range strings.Split(input, ",") {
		value, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		program = append(program, value)
	}

	return program
}

func main() {
	executor.Run(executor.ReadInput("2019/day02.input.txt"), run)
}
