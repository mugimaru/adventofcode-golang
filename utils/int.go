package utils

// SliceInt returns a list of digits of a given integer
//
// SliceInt(123456) => []int8{1,2,3,4,5}
func SliceInt(v int) []int8 {
	var slice []int8

	for {
		if v == 0 {
			break
		}

		i := int8(v % 10)
		v = v / 10
		slice = append([]int8{i}, slice...)
	}

	return slice
}

// SplitIntRange splits range into n sequential ranges
func SplitIntRange(from int, to int, n int) [][]int {
	step := (to - from) / n
	groups := make([][]int, n)

	nextGroupStartsFrom := from
	for i := 0; i < n; i++ {

		groupEndsAt := nextGroupStartsFrom + step
		if i == n-1 {
			if to < groupEndsAt {
				groupEndsAt = to
			}
		}

		groups[i] = []int{nextGroupStartsFrom, groupEndsAt}

		nextGroupStartsFrom = groupEndsAt + 1
	}

	return groups
}
