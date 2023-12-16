package utils

func SumIntSlice(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}
