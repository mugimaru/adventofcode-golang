package grid

// Rotate rotates [][]rune clockwise
func Rotate(grid [][]rune) [][]rune {
	rotated := make([][]rune, len(grid[0]))
	for i := range rotated {
		rotated[i] = make([]rune, len(grid))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			rotated[len(grid[0])-1-j][i] = grid[i][j]
		}
	}

	return rotated
}
