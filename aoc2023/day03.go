package aoc2023

import (
	"strconv"
	"strings"
	"unicode"
)

func init() {
	registerFun("03", SolveDay03)
}

var digits map[rune]int = map[rune]int{
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'0': 0,
}

const dotSym rune = '.'

type point struct {
	i, j int
}

var nullPoint point = point{-1, -1}

func SolveDay03(input string) (interface{}, interface{}) {
	lines := strings.Split(input, "\n")
	var s [][]rune
	for i, line := range lines {
		s = append(s, []rune{})
		for _, char := range line {
			if unicode.IsPrint(char) {
				s[i] = append(s[i], char)
			}
		}
	}

	var currentNumber []rune
	var currentNumbercheckAdjCells bool
	var currentNumberAdjStarPoint point
	var sum int
	gears := make(map[point][]int)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			c := s[i][j]
			_, isDigit := digits[c]

			if !currentNumbercheckAdjCells && isDigit {
				isAdjToSymbol, adjStarPoint := checkAdjCells(s, i, j)
				if isAdjToSymbol {
					currentNumbercheckAdjCells = true
				}

				if adjStarPoint != nullPoint {
					if _, ok := gears[adjStarPoint]; !ok {
						gears[adjStarPoint] = []int{}
					}
					currentNumberAdjStarPoint = adjStarPoint
				}
			}

			if isDigit {
				currentNumber = append(currentNumber, c)
			}

			if len(currentNumber) > 0 && (!isDigit || j == len(s[i])-1) {
				if currentNumbercheckAdjCells {
					num, err := strconv.Atoi(string(currentNumber))
					if err != nil {
						panic(err)
					}
					sum += num

					if currentNumberAdjStarPoint != nullPoint {
						gears[currentNumberAdjStarPoint] = append(gears[currentNumberAdjStarPoint], num)
					}
				}
				currentNumber = []rune{}
				currentNumbercheckAdjCells = false
				currentNumberAdjStarPoint = nullPoint
			}
		}
	}

	var sumOfGearRatios int
	for _, nums := range gears {
		if len(nums) < 2 {
			continue
		}
		gearRatio := nums[0]
		for i := 1; i < len(nums); i++ {
			gearRatio *= nums[i]
		}
		sumOfGearRatios += gearRatio
	}
	return sum, sumOfGearRatios
}

func checkAdjCells(schematics [][]rune, i int, j int) (bool, point) {
	var isAdjToSymbol bool
	starPoint := nullPoint
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			k := i + di
			l := j + dj
			if (k == 0 && l == 0) || k < 0 || l < 0 || k == len(schematics) || l == len(schematics[0]) {
				continue
			}

			if _, isDigit := digits[schematics[k][l]]; !isDigit && schematics[k][l] != dotSym {
				isAdjToSymbol = true
				if schematics[k][l] == '*' {
					starPoint = point{k, l}
				}
			}
		}
	}

	return isAdjToSymbol, starPoint
}
