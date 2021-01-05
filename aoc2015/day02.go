package aoc2015

import (
	"fmt"
	"strings"

	"github.com/mugimaru73/adventofcode-golang/utils"
)

func init() {
	registerFun("02", solveDay02)
}

func solveDay02(input string) (interface{}, interface{}) {
	p1 := 0
	p2 := 0

	for _, line := range strings.Split(input, "\n") {
		var l, w, h int
		if _, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h); err != nil {
			panic(err)
		}

		s1 := l * w
		s2 := w * h
		s3 := h * l

		p1 += 2*s1 + 2*s2 + 2*s3 + utils.MinInt(s1, s2, s3)

		p2 += l * w * h
		if w+l < w+h && w+l < l+h {
			p2 += 2 * (w + l)
		} else if w+h < l+h {
			p2 += 2 * (w + h)
		} else {
			p2 += 2 * (l + h)
		}
	}

	return p1, p2
}
