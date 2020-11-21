package main

import (
	"github.com/mugimaru73/adventofcode-golang/executor"
	"github.com/mugimaru73/adventofcode-golang/intcode"
)

func main() {
	executor.Run(executor.ReadInput("2019/day05.input.txt"), run)
}

func run(input string) (interface{}, interface{}) {
	program := intcode.LoadProgram(input)

	chIn := make(chan int)
	chOut := make(chan int, 100)

	out := []int{}

	go intcode.Run(&program, chIn, chOut)
	chIn <- 1

	for v := range chOut {
		out = append(out, v)
	}

	return out, nil
}
