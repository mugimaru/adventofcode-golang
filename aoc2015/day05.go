package aoc2015

import (
	"strings"
)

func init() {
	registerFun("05", solveDay05)
}

func solveDay05(input string) (interface{}, interface{}) {
	var niceStrings, niceStrings2 int
	for _, str := range strings.Split(input, "\n") {
		if isStringNice(str) {
			niceStrings++
		}

		if isStringNice2(str) {
			niceStrings2++
		}
	}

	return niceStrings, niceStrings2
}

func isStringNice2(str string) bool {
	rs := []rune(str)
	var hasTriplet, hasDoublePair bool
	pairs := make(map[string][]int)

	for i := range rs {
		if i < len(rs)-2 && rs[i] == rs[i+2] {
			hasTriplet = true
		}

		if i < len(rs)-1 {
			pairs[string(rs[i:i+2])] = append(pairs[string(rs[i:i+2])], i)
		}
	}

	for _, ind := range pairs {
		if len(ind) <= 1 {
			continue
		}

		if len(ind) > 2 || (len(ind) == 2 && ind[1]-ind[0] > 1) {
			hasDoublePair = true
			break
		}
	}

	return hasTriplet && hasDoublePair
}

func isStringNice(str string) bool {
	vowels := 0
	hasDoubleLetter := false

	for i, char := range str {
		if strings.Contains("aeiou", string(char)) {
			vowels++
		}

		if i != len(str)-1 {
			pair := string(str[i : i+2])
			if pair == "ab" || pair == "cd" || pair == "pq" || pair == "xy" {
				return false
			}

			if str[i] == str[i+1] {
				hasDoubleLetter = true
			}
		}
	}

	return hasDoubleLetter && vowels >= 3
}
