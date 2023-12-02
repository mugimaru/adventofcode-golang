package aoc2023

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func init() {
	registerFun("02", SolveDay02)
}

type cubesSet struct {
	R int
	G int
	B int
}

func (cs *cubesSet) Add(cs2 *cubesSet) {
	cs.R += cs2.R
	cs.G += cs2.G
	cs.B += cs2.B
}

type round struct {
	id   int
	sets []*cubesSet
}

var bag *cubesSet = &cubesSet{R: 12, G: 13, B: 14}

func SolveDay02(input string) (interface{}, interface{}) {
	rounds := parseDay02Input(input)

	return solveDay02Part1(rounds), 0
}

func solveDay02Part1(rounds []*round) int {
	var sum int
	fmt.Println(bag)
	for i := 0; i < len(rounds); i++ {
		isGameValid := true
		for _, s := range rounds[i].sets {
			if bag.R < s.R || bag.G < s.G || bag.B < s.B {
				isGameValid = false
				break
			}
		}

		if isGameValid {
			sum += rounds[i].id
		}
	}

	return sum
}

func parseDay02Input(input string) []*round {
	var err error
	lines := strings.Split(input, "\n")
	rounds := make([]*round, len(lines))

	for i := 0; i < len(lines); i++ {
		rounds[i] = &round{}

		gameAndRounds := strings.Split(lines[i], ": ")

		rounds[i].id, err = strconv.Atoi(strings.Fields(gameAndRounds[0])[1])
		if err != nil {
			panic(err)
		}

		setsStr := strings.Split(gameAndRounds[1], "; ")
		rounds[i].sets = make([]*cubesSet, len(setsStr))
		for si := 0; si < len(setsStr); si++ {
			rounds[i].sets[si] = &cubesSet{}

			cubes := strings.Split(setsStr[si], ", ")
			for ci := 0; ci < len(cubes); ci++ {
				amountAndColor := strings.Fields(cubes[ci])

				value, err := strconv.Atoi(amountAndColor[0])
				if err != nil {
					panic(err)
				}
				switch amountAndColor[1] {
				case "red":
					rounds[i].sets[si].R += value
				case "green":
					rounds[i].sets[si].G += value
				case "blue":
					rounds[i].sets[si].B += value
				default:
					log.Fatalf("unexpected color %s", amountAndColor[1])
				}
			}
		}
	}

	return rounds
}
