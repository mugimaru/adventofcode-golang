package aoc2019

import "github.com/mugimaru73/adventofcode-golang/intcode"

func init() {
	registerFun("09", SolveDay09)
}

func SolveDay09(input string) (interface{}, interface{}) {
	mem := intcode.LoadProgram(input)
	mem2 := intcode.CopyProgram(mem)

	return intcode.RunAndRead(&mem, []int64{1}), intcode.RunAndRead(&mem2, []int64{2})
}
