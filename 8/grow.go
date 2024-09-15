package kata8
/*
#Beginner - Reduce but Grow

Given a non-empty array of integers, return the result of multiplying the values together in order. Example:

[1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24
*/
func Grow(arr []int) int {
	result := 1
	for _, value := range arr {
		result *= value
	}
	return result
}