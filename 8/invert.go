package kata8

/*
#Invert values

Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []
You can assume that all values are integers. Do not mutate the input array.
*/
func Invert(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = -v
	}
	return result
}

func Invert1(arr []int) (out []int) {
	for _,el := range arr { out = append(out,-el) }
	return
}