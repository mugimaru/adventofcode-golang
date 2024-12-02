package utils

func SumIntSlice(s []int) (sum int) {
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return sum
}
func RemoveElementByIndexInplace(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}
func RemoveElementByIndex(s []int, i int) []int {
	res := make([]int, len(s))
	copy(res, s)
	
	return RemoveElementByIndexInplace(res, i)
}
