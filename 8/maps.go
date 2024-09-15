package kata8

/*
#Beginner - Lost Without a Map

Given an array of integers, return a new array with each value doubled.

For example:

[1, 2, 3] --> [2, 4, 6]
*/
func Maps(x []int) []int {
	var result = make([]int, len(x))
	for index, value := range x {
		result[index] = value * 2
	}
	return result
}

func Maps2(x []int)  (result []int) {
	for _,j := range x {
	  result = append(result,j*2)
	}
	return
}