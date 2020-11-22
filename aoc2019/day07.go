package aoc2019

import (
	"github.com/mugimaru73/adventofcode-golang/intcode"
	"github.com/mugimaru73/adventofcode-golang/utils"
)

func init() {
	registerFun("07", SolveDay07)
}

const n = 5

func SolveDay07(input string) (interface{}, interface{}) {
	p := intcode.LoadProgram(input)
	return maxThrust(p, []int{0, 1, 2, 3, 4}, testSettings), maxThrust(p, []int{5, 6, 7, 8, 9}, testSettingsFeedbackLoop)
}

func maxThrust(program intcode.Memory, settingsTemplate []int, fun func([5]int, intcode.Memory) int) int {
	max := 0
	for _, settings := range utils.Permutations(settingsTemplate) {
		var arr [5]int
		copy(arr[:], settings)

		signal := fun(arr, program)
		if signal > max {
			max = signal
		}
	}

	return max
}

func testSettings(settings [5]int, program intcode.Memory) int {
	out := 0

	for i := 0; i < 5; i++ {
		p := intcode.CopyProgram(program)
		chIn := make(chan int64)
		chOut := make(chan int64)
		go intcode.Run(&p, chIn, chOut, nil)
		chIn <- int64(settings[i])
		chIn <- int64(out)

		out = int(<-chOut)
	}

	return out
}

func testSettingsFeedbackLoop(settings [5]int, program intcode.Memory) int {
	ch := make([](chan int64), n)
	for i := 0; i < n; i++ {
		ch[i] = make(chan int64, 100)
		ch[i] <- int64(settings[i])
	}
	chDone := make(chan int)

	for i := 0; i < n; i++ {
		p := intcode.CopyProgram(program)
		if i == 0 {
			ch[n-1] <- 0
			go intcode.Run(&p, ch[n-1], ch[0], nil)
		} else if i == n-1 {
			go intcode.Run(&p, ch[n-2], ch[i], chDone)
		} else {
			go intcode.Run(&p, ch[i-1], ch[i], nil)
		}
	}

	<-chDone
	return int(<-ch[n-1])
}
