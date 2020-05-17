package main

import (
	"strconv"
	"strings"

	"github.com/mugimaru73/adventofcode-golang/executor"
	"github.com/mugimaru73/adventofcode-golang/utils"
)

var concurrency = 4

func run(input string) (interface{}, interface{}) {
	pwdRange := []int{0, 0}

	for i, v := range strings.Split(input, "-") {
		intValue, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		pwdRange[i] = intValue

	}

	return countValidPasswords(pwdRange[0], pwdRange[1], isPwdValid), countValidPasswords(pwdRange[0], pwdRange[1], isPwdValidV2)
}

func countValidPasswords(from int, to int, validator func(int) bool) int {
	count := 0
	countersChan := make(chan int, concurrency)

	for _, r := range utils.SplitIntRange(from, to, concurrency) {
		go doCountValidPasswords(r[0], r[1], validator, countersChan)
	}

	for a := 0; a < concurrency; a++ {
		count += <-countersChan
	}

	return count
}

func doCountValidPasswords(from int, to int, validator func(int) bool, res chan<- int) {
	var validPasswords []int

	for pwd := from; pwd <= to; pwd++ {
		if validator(pwd) {
			validPasswords = append(validPasswords, pwd)
		}
	}

	res <- len(validPasswords)
}

func isPwdValid(pwd int) bool {
	digits := utils.SliceInt(pwd)
	hasDoubleDigit := false

	for i := 0; i < len(digits)-1; i++ {
		switch {
		case digits[i] == digits[i+1]:
			hasDoubleDigit = true
		case digits[i] > digits[i+1]:
			return false
		}
	}

	return hasDoubleDigit
}

func isPwdValidV2(pwd int) bool {
	digits := utils.SliceInt(pwd)
	hasDoubleDigit := false

	seqDigit := digits[0]
	seqLen := 1

	for i := 1; i < len(digits); i++ {
		if digits[i] < digits[i-1] {
			return false
		}

		if digits[i] == seqDigit && i != len(digits)-1 {
			seqLen++
		} else {
			if i == len(digits)-1 {
				seqLen++
			}

			if seqLen == 2 {
				hasDoubleDigit = true
			}
			seqDigit = digits[i]
			seqLen = 1
		}
	}

	return hasDoubleDigit
}

func main() {
	executor.Run(executor.ReadInput("2019/day04.input.txt"), run)
}
