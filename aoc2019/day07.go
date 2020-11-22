package aoc2019

import (
	"github.com/mugimaru73/adventofcode-golang/intcode"
	"github.com/mugimaru73/adventofcode-golang/utils"
)

func init() {
	registerFun("07", SolveDay07)
}

func SolveDay07(input string) (interface{}, interface{}) {
	program := intcode.LoadProgram(input)

	max := 0
	for _, settings := range utils.Permutations([]int{0, 1, 2, 3, 4}) {
		var arr [5]int
		copy(arr[:], settings)

		signal := testSettings(arr, program)
		if signal > max {
			max = signal
		}
	}

	return max, nil
}

func testSettings(settings [5]int, program []int) int {
	out := 0

	for i := 0; i < 5; i++ {
		p := intcode.CopyProgram(program)
		chIn := make(chan int)
		chOut := make(chan int)
		go intcode.Run(&p, chIn, chOut)
		chIn <- settings[i]
		chIn <- out

		close(chIn)
		out = <-chOut
	}

	return out
}
