package aoc2019

import (
	"strings"
)

func init() {
	registerFun("08", SolveDay08)
}

const w = 25
const h = 6

func SolveDay08(strInput string) (interface{}, interface{}) {
	input := []byte(strInput)
	size := w * h
	layers := len(input) / size

	digits := make([][]int, layers)
	for i := 0; i < layers; i++ {
		digits[i] = make([]int, 3)
	}
	for i := 0; i < len(input); i++ {
		layer := i / size
		digits[layer][input[i]-'0']++
	}

	p1 := digits[0]
	for i := 1; i < layers; i++ {
		if digits[i][0] < p1[0] {
			p1 = digits[i]
		}
	}

	builder := &strings.Builder{}
	builder.WriteByte('\n')
	for ih := 0; ih < h; ih++ {
		for iw := 0; iw < w; iw++ {
			for il := 0; il < layers; il++ {
				v := input[il*size+ih*w+iw]
				if v == '0' {
					builder.WriteByte(' ')
					break
				}

				if v == '1' {
					builder.WriteByte('X')
					break
				}
			}
		}
		builder.WriteString("\n")
	}

	return p1[1] * p1[2], builder.String()
}
