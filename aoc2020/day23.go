package aoc2020

import (
	"fmt"
	"strconv"
)

func init() {
	registerFun("23", SolveDay23)
}

type crabCupsGame struct {
	current  *crabCup
	maxLabel int
	minLabel int
	nodes    map[int]*crabCup
}

type crabCup struct {
	label int
	next  *crabCup
	prev  *crabCup
}

func newCrabCupsGame(input string, part2 bool) crabCupsGame {
	game := crabCupsGame{nodes: make(map[int]*crabCup)}

	var prevCup *crabCup
	i := 0

	for {
		var label int

		if i < len(input) {
			label, _ = strconv.Atoi(string(input[i]))
		} else {
			if !part2 {
				break
			}

			if game.maxLabel == 1000000 {
				break
			}

			label = game.maxLabel + 1
		}

		if label > game.maxLabel || i == 0 {
			game.maxLabel = label
		}

		if label < game.minLabel || i == 0 {
			game.minLabel = label
		}

		newCup := &crabCup{label: label, next: game.current}
		game.nodes[label] = newCup

		if game.current == nil {
			game.current = newCup
		} else {
			newCup.prev = prevCup
			prevCup.next = newCup
			game.current.prev = newCup
		}

		prevCup = newCup
		i++
	}

	return game
}

func (g *crabCupsGame) move() {
	pickedCups := g.pickThreeCups()
	destCupLabel := g.selectDestinationCupLabel(pickedCups)

	g.insertCups(destCupLabel, pickedCups)
	g.rotateCurrentCup()
}

func (g *crabCupsGame) pickThreeCups() [3]*crabCup {
	c1 := g.current.next
	c2 := c1.next
	c3 := c2.next

	g.current.next = c3.next
	g.current.next.prev = g.current

	return [3]*crabCup{c1, c2, c3}
}

func (g *crabCupsGame) rotateCurrentCup() {
	g.current = g.current.next
}

func (g *crabCupsGame) insertCups(destLabel int, pickedCups [3]*crabCup) {
	destCup := g.nodes[destLabel]
	nextCup := destCup.next

	pickedCups[0].prev = destCup
	destCup.next = pickedCups[0]

	pickedCups[2].next = nextCup
	nextCup.prev = pickedCups[2]
}

func (g *crabCupsGame) selectDestinationCupLabel(pickedCups [3]*crabCup) int {
	label := g.current.label - 1
	for {
		labelOk := true

		if label < g.minLabel {
			label = g.maxLabel
		}

		for _, pc := range pickedCups {
			if label == pc.label {
				labelOk = false
				label--
				break
			}
		}

		if labelOk {
			break
		}
	}

	return label
}

func (g crabCupsGame) collectCups(afterLabel int) (res string) {
	firstNode := g.nodes[afterLabel]
	node := firstNode.next
	for node != firstNode {
		res += fmt.Sprint(node.label)
		node = node.next
	}

	return
}

func SolveDay23(input string) (interface{}, interface{}) {
	game := newCrabCupsGame(input, false)

	for i := 0; i < 100; i++ {
		game.move()
	}

	game2 := newCrabCupsGame(input, true)
	for i := 0; i < 10000000; i++ {
		game2.move()
	}

	return game.collectCups(1), game2.nodes[1].next.label * game2.nodes[1].next.next.label
}
