package aoc2019

import (
	"github.com/mugimaru73/adventofcode-golang/intcode"
)

func init() {
	registerFun("05", SolveDay05)
}

func SolveDay05(input string) (interface{}, interface{}) {
	program := intcode.LoadProgram(input)

	p2 := intcode.CopyProgram(program)
	return intcode.RunAndRead(&program, []int64{1}), intcode.RunAndRead(&p2, []int64{5})
}
