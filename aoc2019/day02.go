package aoc2019

import (
	"github.com/mugimaru73/adventofcode-golang/intcode"
)

const p2DesiredOutput = 19690720

func init() {
	registerFun("02", SolveDay02)
}

func SolveDay02(input string) (interface{}, interface{}) {
	program := intcode.LoadProgram(input)
	return calculateOutput(12, 2, program), solveP2(program)
}

func solveP2(program intcode.Memory) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			if calculateOutput(noun, verb, program) == p2DesiredOutput {
				return 100*noun + verb
			}
		}
	}

	panic("Solution not found")
}

func calculateOutput(noun int, verb int, program intcode.Memory) int {
	mem := intcode.CopyProgram(program)
	mem.Set(1, intcode.Value(noun))
	mem.Set(2, intcode.Value(verb))
	intcode.Run(&mem, nil, nil, nil)

	return int(mem.Get(0))
}
