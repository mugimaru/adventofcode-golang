package aoc2020

import "strings"

func init() {
	registerFun("24", SolveDay24)
}

type point struct {
	x int
	y int
}

func SolveDay24(input string) (interface{}, interface{}) {
	blackTiles := make(map[point]bool)
	for _, line := range strings.Split(input, "\n") {
		tile := parseTile(line)
		if blackTiles[tile] {
			delete(blackTiles, tile)
		} else {
			blackTiles[tile] = true
		}
	}

	p1 := len(blackTiles)

	for i := 0; i < 100; i++ {
		blackTiles = dailyFlip(blackTiles)
	}

	return p1, len(blackTiles)
}

func parseTile(line string) (p point) {
	for len(line) > 0 {
		var step string
		if line[0] == 's' || line[0] == 'n' {
			step, line = line[0:2], line[2:]
		} else {
			step, line = line[0:1], line[1:]
		}

		switch step {
		case "e":
			p.x++
		case "w":
			p.x--
		case "se":
			p.y--
		case "sw":
			p.x--
			p.y--
		case "ne":
			p.x++
			p.y++
		case "nw":
			p.y++
		}
	}

	return
}

func dailyFlip(current map[point]bool) map[point]bool {
	toCheck := make(map[point]bool)
	for k := range current {
		toCheck[k] = true
		for _, n := range getNeighbors(k) {
			toCheck[n] = true
		}
	}

	next := make(map[point]bool)
	for k := range toCheck {
		isBlack := current[k]
		blackNeighbors := 0
		for _, n := range getNeighbors(k) {
			if current[n] {
				blackNeighbors++
			}
		}
		if (isBlack && (blackNeighbors == 1 || blackNeighbors == 2)) || (!isBlack && blackNeighbors == 2) {
			next[k] = true
		}
	}

	return next
}

func getNeighbors(c point) (res []point) {
	return []point{
		{c.x + 1, c.y},
		{c.x - 1, c.y},
		{c.x, c.y - 1},
		{c.x - 1, c.y - 1},
		{c.x + 1, c.y + 1},
		{c.x, c.y + 1},
	}
}
