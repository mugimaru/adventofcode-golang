package aoc2020

import (
	"strings"
)

func init() {
	registerFun("06", SolveDay06)
}

func SolveDay06(input string) (interface{}, interface{}) {
	p1 := 0
	p2 := 0

	for _, group := range strings.Split(input, "\n\n") {
		answers := make(map[rune]int)
		people := 1

		for _, q := range []rune(group) {
			if q == '\n' {
				people++
			} else {
				answers[q]++
			}
		}

		for _, v := range answers {
			p1++
			if v == people {
				p2++
			}
		}
	}

	return p1, p2
}
