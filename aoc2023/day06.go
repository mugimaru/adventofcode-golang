package aoc2023

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registerFun("06", SolveDay06)
}

func SolveDay06(input string) (interface{}, interface{}) {
	p1Input := parseDay06Input(input)
	var p1 int
	for i := 0; i < len(p1Input[0]); i++ {
		time := p1Input[0][i]
		distance := p1Input[1][i]
		res := numberOfWaysToWin(time, distance)
		if p1 == 0 {
			p1 = res
		} else {
			p1 *= res
		}
	}
	p2Input := parseDay06Input(strings.ReplaceAll(input, " ", ""))
	return p1, numberOfWaysToWin(p2Input[0][0], p2Input[1][0])
}

func numberOfWaysToWin(time, distance int) int {
	var res int

	for speed := 1; speed < time-1; speed++ {
		timeLeft := time - speed
		if timeLeft*speed > distance {
			res++
		}
	}

	return res
}

func parseDay06Input(input string) [2][]int {
	re := regexp.MustCompile("[0-9]+")
	var inputs [2][]int

	for i, s := range strings.Split(input, "\n") {
		nums := re.FindAllStringSubmatch(s, -1)
		inputs[i] = make([]int, len(nums))
		for j, n := range nums {
			num, err := strconv.Atoi(n[0])
			if err != nil {
				panic(err)
			}

			inputs[i][j] = num
		}
	}

	return inputs
}
