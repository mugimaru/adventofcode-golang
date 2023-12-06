package utils

import (
	"strconv"
	"strings"
)

func ReverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func ParseNumsRow(input string) []int {
	items := strings.Split(input, " ")
	nums := make([]int, len(items))
	for i := 0; i < len(items); i++ {
		num, err := strconv.Atoi(items[i])
		if err != nil {
			panic(err)
		}
		nums[i] = num
	}
	return nums
}
