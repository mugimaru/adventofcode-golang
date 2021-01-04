package grid

import (
	"fmt"
	"strings"
)

// ParseGrid parses string into [][]rune grid
func ParseGrid(input string) (grid [][]rune) {
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []rune(row))
	}

	return
}

// Print prints [][]rune to stdout
func Print(grid [][]rune) {
	fmt.Println(Sprint(grid))
}

// Sprint prints [][]rune to string
func Sprint(grid [][]rune) (out string) {
	for _, row := range grid {
		out += fmt.Sprintf("\n%s", string(row))
	}
	return
}

// Equal checks if two [][]rune grids are equal
func Equal(g1 [][]rune, g2 [][]rune) bool {
	if len(g1) != len(g2) {
		return false
	}

	for i := 0; i < len(g1); i++ {
		if len(g1[i]) != len(g2[i]) {
			return false
		}

		for j := 0; j < len(g1[i]); j++ {
			if g1[i][j] != g2[i][j] {
				return false
			}
		}
	}

	return true
}
