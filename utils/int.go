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
