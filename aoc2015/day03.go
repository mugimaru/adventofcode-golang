package aoc2015

func init() {
	registerFun("03", solveDay03)
}

type point struct {
	x int
	y int
}

func solveDay03(input string) (interface{}, interface{}) {
	visits1 := make(map[point]bool)

	pos := point{0, 0}
	visits1[pos] = true
	for _, dir := range input {
		switch dir {
		case '^':
			pos.y++
		case 'v':
			pos.y--
		case '>':
			pos.x++
		case '<':
			pos.x--
		}
		visits1[pos] = true
	}

	visits2 := make(map[point]bool)
	santa := point{0, 0}
	roboSanta := point{0, 0}

	moves := &santa
	visits2[*moves] = true

	for _, dir := range input {
		switch dir {
		case '^':
			moves.y++
		case 'v':
			moves.y--
		case '>':
			moves.x++
		case '<':
			moves.x--
		}
		visits2[*moves] = true

		if moves == &santa {
			moves = &roboSanta
		} else {
			moves = &santa
		}
	}

	return len(visits1), len(visits2)
}
