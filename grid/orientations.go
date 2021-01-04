package grid

// AllOrientations all possible orientations of [][]rune grid
func AllOrientations(grid [][]rune) [][][]rune {
	res := [][][]rune{grid}

	for i := 0; i < 3; i++ {
		res = append(res, Rotate(res[i]))
	}

	for i := 0; i < 4; i++ {
		res = append(res, Flip(res[i]))
	}

	return res
}
