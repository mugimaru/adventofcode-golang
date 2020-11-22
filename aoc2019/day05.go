package aoc2019

import (
	"github.com/mugimaru73/adventofcode-golang/intcode"
)

func init() {
	registerFun("05", SolveDay05)
}

func SolveDay05(input string) (interface{}, interface{}) {
	program := intcode.LoadProgram(input)
	outCh1 := doRunProgram(program, []int{1})
	outCh2 := doRunProgram(program, []int{5})

	return readOutput(outCh1), readOutput(outCh2)
}

func doRunProgram(p []int, input []int) chan int {
	program := intcode.CopyProgram(p)
	chIn := make(chan int)
	chOut := make(chan int, 100)

	go intcode.Run(&program, chIn, chOut)
	for _, v := range input {
		chIn <- v
	}

	return chOut
}

func readOutput(chOut chan int) []int {
	out := []int{}
	for v := range chOut {
		out = append(out, v)
	}
	return out
}
