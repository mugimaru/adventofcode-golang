package aoc2020

import (
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registerFun("02", SolveDay02)
}

func SolveDay02(input string) (interface{}, interface{}) {
	countPart1 := 0
	countPart2 := 0

	var lineRegExp = regexp.MustCompile(`(?P<min>\d+)-(?P<max>\d+) (?P<letter>[a-z]): (?P<password>.+)`)

	for _, line := range strings.Split(input, "\n") {
		match := lineRegExp.FindStringSubmatch(line)

		i1, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		i2, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		letter := match[3]
		password := []rune(match[4])

		matchesP1 := 0
		for _, v := range password {
			if string(v) == letter {
				matchesP1++
			}
		}
		if matchesP1 >= i1 && matchesP1 <= i2 {
			countPart1++
		}

		matchesP2 := 0
		if string(password[i1-1]) == letter {
			matchesP2++
		}
		if string(password[i2-1]) == letter {
			matchesP2++
		}

		if matchesP2 == 1 {
			countPart2++
		}

	}

	return countPart1, countPart2
}
