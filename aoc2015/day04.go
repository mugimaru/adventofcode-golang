package aoc2015

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func init() {
	registerFun("04", solveDay04)
}

func solveDay04(input string) (interface{}, interface{}) {
	var p1, p2 int

	for i := 0; ; i++ {
		h := md5.New()
		io.WriteString(h, input)
		io.WriteString(h, fmt.Sprint(i))
		hash := fmt.Sprintf("%x", h.Sum(nil))

		if strings.HasPrefix(hash, "00000") && p1 <= 0 {
			p1 = i
		}

		if strings.HasPrefix(hash, "000000") && p2 <= 0 {
			p2 = i
		}

		if p1 > 0 && p2 > 0 {
			break
		}
	}

	return p1, p2
}
