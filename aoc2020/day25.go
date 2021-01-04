package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("25", SolveDay25)
}

func findCardLoopSize(pk int) (ls int) {
	acc := 1
	for acc != pk {
		ls++
		acc = (acc * 7) % 20201227
	}

	return
}

func findDoorLoopSize(pk int, cardLS int) int {
	doorLS := 1
	for i := 0; i < cardLS; i++ {
		doorLS = (doorLS * pk) % 20201227
	}
	return doorLS
}

func SolveDay25(input string) (interface{}, interface{}) {
	inp := strings.Split(input, "\n")
	cardPK, _ := strconv.Atoi(inp[0])
	doorPK, _ := strconv.Atoi(inp[1])

	return findDoorLoopSize(doorPK, findCardLoopSize(cardPK)), nil
}
