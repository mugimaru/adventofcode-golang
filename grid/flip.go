package grid

// Flip flips (mirrors) [][]rune grid
func Flip(grid [][]rune) (flipped [][]rune) {
	for i := range grid {
		flipped = append(flipped, []rune{})
		for j := len(grid[i]) - 1; j >= 0; j-- {
			flipped[i] = append(flipped[i], grid[i][j])
		}
	}
	return flipped
}
