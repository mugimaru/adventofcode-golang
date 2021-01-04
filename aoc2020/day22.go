package aoc2020

import (
	"strconv"
	"strings"
)

func init() {
	registerFun("22", SolveDay22)
}

func areDecksEqual(d1 []int, d2 []int) bool {
	if len(d1) != len(d2) {
		return false
	}
	for i := 0; i < len(d1); i++ {
		if d1[i] != d2[i] {
			return false
		}
	}

	return true
}

func SolveDay22(input string) (interface{}, interface{}) {
	players := [][]int{}

	for _, lines := range strings.Split(input, "\n\n") {
		playerDeck := []int{}
		for _, v := range strings.Split(lines, "\n")[1:] {
			card, _ := strconv.Atoi(v)
			playerDeck = append(playerDeck, card)
		}

		players = append(players, playerDeck)
	}

	winner1 := playCombat(players[0], players[1])
	winner2, _ := playRecursiveCombat(players[0], players[1])

	return combatScore(winner1), combatScore(winner2)
}

func combatScore(winner []int) (score int) {
	mult := 1
	for i := len(winner) - 1; i >= 0; i-- {
		score += winner[i] * mult
		mult++
	}

	return
}

func playRecursiveCombat(p1 []int, p2 []int) ([]int, bool) {
	history := [][2][]int{}

	for {
		for i := 0; i < len(history); i++ {
			if areDecksEqual(p1, history[i][0]) || areDecksEqual(p2, history[i][1]) {
				return p1, true
			}
		}

		history = append(history, [2][]int{p1, p2})

		var (
			p1c    int
			p2c    int
			p1Wins bool
		)
		p1, p1c = p1[1:], p1[0]
		p2, p2c = p2[1:], p2[0]

		if p1c <= len(p1) && p2c <= len(p2) {
			newDeck1, newDeck2 := make([]int, p1c), make([]int, p2c)

			for i := 0; i < p1c; i++ {
				newDeck1[i] = p1[i]
			}
			for i := 0; i < p2c; i++ {
				newDeck2[i] = p2[i]
			}

			_, p1Wins = playRecursiveCombat(newDeck1, newDeck2)
		} else {
			p1Wins = p1c > p2c
		}

		if p1Wins {
			p1 = append(p1, p1c, p2c)
		} else {
			p2 = append(p2, p2c, p1c)
		}

		if len(p1) == 0 {
			return p2, false
		}

		if len(p2) == 0 {
			return p1, true
		}
	}
}

func playCombat(p1 []int, p2 []int) []int {
	for {
		if len(p1) == 0 {
			return p2
		}

		if len(p2) == 0 {
			return p1
		}

		var (
			p1c int
			p2c int
		)

		p1, p1c = p1[1:], p1[0]
		p2, p2c = p2[1:], p2[0]
		if p2c > p1c {
			p2 = append(p2, p2c, p1c)
		} else {
			p1 = append(p1, p1c, p2c)
		}
	}
}
