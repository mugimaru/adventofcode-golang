package aoc2020

import (
	"fmt"
	"math"
	"strings"

	"github.com/mugimaru73/adventofcode-golang/grid"
)

func init() {
	registerFun("20", SolveDay20)
}

type tile struct {
	id   int
	grid [][]rune
}

type monsterPattern struct {
	pattern [][]int
	width   int
}

func SolveDay20(input string) (interface{}, interface{}) {
	tiles := []tile{}
	for _, t := range strings.Split(input, "\n\n") {
		lines := strings.Split(t, "\n")
		var id int
		if _, err := fmt.Sscanf(lines[0], "Tile %d:", &id); err != nil {
			panic(err)
		}

		t := make([][]rune, len(lines)-1)
		for i, row := range lines[1:] {
			t[i] = []rune(row)
		}

		tiles = append(tiles, tile{id, t})
	}

	s := int(math.Sqrt(float64(len(tiles))))

	pic := make([][]tile, s)
	for i := 0; i < s; i++ {
		pic[i] = make([]tile, s)
	}

	pic, _ = assemblePicture(pic, tiles, make(map[int]bool))
	p1 := pic[0][0].id * pic[0][s-1].id * pic[s-1][0].id * pic[s-1][s-1].id

	flatPic := removeBordersAndFlatten(pic)

	mPattern := parseMonsterPattern([]string{"                  # ", "#    ##    ##    ###", " #  #  #  #  #  #"})

	roughness := 0
	for _, orient := range grid.AllOrientations(flatPic) {
		coords := findMonsters(orient, mPattern)
		if len(coords) > 0 {
			orient = markMonsters(orient, coords, mPattern)

			for _, row := range orient {
				for _, char := range row {
					if char == '#' {
						roughness++
					}
				}
			}

			break
		}
	}

	return p1, roughness
}

func markMonsters(pic [][]rune, coords [][2]int, mp monsterPattern) [][]rune {
	for _, point := range coords {
		row := point[0]
		col := point[1]

		for mpRowI, mpRow := range mp.pattern {
			for _, mpColOff := range mpRow {
				pic[row+mpRowI][col+mpColOff] = 'O'
			}
		}
	}

	return pic
}

func findMonsters(pic [][]rune, mp monsterPattern) (monstersCoord [][2]int) {
	for row := 0; row <= len(pic)-len(mp.pattern); row++ {
		for col := 0; col <= len(pic)-mp.width; col++ {
			if hasMonster(pic, mp, row, col) {
				monstersCoord = append(monstersCoord, [2]int{row, col})
			}
		}
	}

	return
}

func hasMonster(pic [][]rune, mp monsterPattern, row int, col int) bool {
	for rowI, rowPattern := range mp.pattern {
		for _, offset := range rowPattern {
			if pic[row+rowI][col+offset] != '#' {
				return false
			}
		}
	}

	return true
}

func parseMonsterPattern(lines []string) monsterPattern {
	w := len(lines[0])
	pattern := make([][]int, len(lines))

	for li, line := range lines {
		lp := []int{}
		for i, v := range line {
			if v == '#' {
				lp = append(lp, i)
			}
		}
		pattern[li] = lp
	}

	return monsterPattern{width: w, pattern: pattern}
}

func removeBordersAndFlatten(pic [][]tile) [][]rune {
	flatPicLen := len(pic) * (len(pic[0][0].grid) - 2)
	flatPic := make([][]rune, flatPicLen)
	for i := 0; i < flatPicLen; i++ {
		flatPic[i] = make([]rune, flatPicLen)
	}

	for i := 0; i < len(pic); i++ {
		for j := 0; j < len(pic); j++ {
			g := grid.RemoveBorders(pic[i][j].grid)
			gLen := len(g)
			for k := 0; k < len(g); k++ {
				for l := 0; l < len(g); l++ {
					flatPic[gLen*i+k][gLen*j+l] = g[k][l]
				}
			}
		}
	}

	return flatPic
}

func assemblePicture(pic [][]tile, tiles []tile, used map[int]bool) ([][]tile, bool) {
	for row := 0; row < len(pic); row++ {
		for col := 0; col < len(pic); col++ {
			if pic[row][col].id != 0 {
				continue
			}

			for _, t := range tiles {
				if used[t.id] {
					continue
				}

				for _, orient := range grid.AllOrientations(t.grid) {
					if row != 0 {
						topRow := grid.GetRow(orient, 0)
						bottomRow := grid.GetRow(pic[row-1][col].grid, len(t.grid)-1)
						if !grid.BordersEqual(topRow, bottomRow) {
							continue
						}

					}
					if col != 0 {
						leftCol := grid.GetCol(orient, 0)
						rightCol := grid.GetCol(pic[row][col-1].grid, len(t.grid)-1)
						if !grid.BordersEqual(leftCol, rightCol) {
							continue
						}
					}

					pic[row][col] = tile{id: t.id, grid: orient}
					used[t.id] = true
					if res, ok := assemblePicture(pic, tiles, used); ok {
						return res, ok
					}
					pic[row][col] = tile{}
					used[t.id] = false
				}
			}

			if pic[row][col].id == 0 {
				return pic, false
			}
		}
	}

	return pic, true
}
