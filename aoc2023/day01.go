package aoc2023

import (
	"fmt"
	"github.com/mugimaru73/adventofcode-golang/utils"
	"regexp"
	"strconv"
	"strings"
)

var stringsToDigits = map[string]int{
	"one":   1,
	"1":     1,
	"two":   2,
	"2":     2,
	"three": 3,
	"3":     3,
	"four":  4,
	"4":     4,
	"five":  5,
	"5":     5,
	"six":   6,
	"6":     6,
	"seven": 7,
	"7":     7,
	"eight": 8,
	"8":     8,
	"nine":  9,
	"9":     9,
}
var digitsRegexpPart1 *regexp.Regexp = regexp.MustCompile("[0-9]")
var digitsRegexpPart2 *regexp.Regexp = regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|[0-9]")
var digitsRegexpRevPart2 *regexp.Regexp = regexp.MustCompile("enin|thgie|neves|xis|evif|ruof|eerht|owt|eno|[0-9]")

func init() {
	registerFun("01", SolveDay01)
}

func SolveDay01(input string) (interface{}, interface{}) {
	var sumP1, sumP2 int
	parsedInput := strings.Split(input, "\n")

	for _, line := range parsedInput {
		sumP1 += findNumberDay01(line, digitsRegexpPart1, digitsRegexpPart1)
		sumP2 += findNumberDay01(line, digitsRegexpPart2, digitsRegexpRevPart2)
	}
	return sumP1, sumP2
}

func findNumberDay01(line string, re *regexp.Regexp, revRe *regexp.Regexp) int {
	firstMatch := re.FindString(line)
	lastMatch := utils.ReverseString(revRe.FindString(utils.ReverseString(line)))

	value, err := strconv.Atoi(fmt.Sprintf("%d%d", stringsToDigits[firstMatch], stringsToDigits[lastMatch]))
	if err != nil {
		panic(err)
	}
	return value
}
