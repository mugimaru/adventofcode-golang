package grid

// GetRow returns n'th row of [][]rune
func GetRow(grid [][]rune, n int) []rune {
	return grid[n]
}

// GetCol returns n'th column of [][]rune
func GetCol(grid [][]rune, n int) []rune {
	col := make([]rune, len(grid))

	for i := 0; i < len(grid); i++ {
		col[i] = grid[i][n]
	}

	return col
}

// BordersEqual tests []rune equality
func BordersEqual(b1 []rune, b2 []rune) bool {
	if len(b1) != len(b2) {
		return false
	}

	for i := 0; i < len(b1); i++ {
		if b1[i] != b2[i] {
			return false
		}
	}

	return true
}

// RemoveBorders removes borders
func RemoveBorders(grid [][]rune) [][]rune {
	ng := make([][]rune, len(grid)-2)

	for i := 0; i < len(ng); i++ {
		ng[i] = make([]rune, len(grid[i+1])-2)
		for j := 0; j < len(ng[i]); j++ {
			ng[i][j] = grid[i+1][j+1]
		}
	}

	return ng
}
