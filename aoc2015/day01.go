package aoc2015

func init() {
	registerFun("01", solveDay01)
}

func solveDay01(input string) (interface{}, interface{}) {
	p1 := 0
	p2 := 0
	for i, char := range input {
		switch char {
		case '(':
			p1++
		case ')':
			p1--
		}

		if p1 < 0 && p2 <= 0 {
			p2 = i + 1
		}
	}

	return p1, p2
}
