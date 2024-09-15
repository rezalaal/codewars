package kata7

/*
#Simple remove duplicates

Remove the duplicates from a list of integers, keeping the last ( rightmost ) occurrence of each element.

Example:
For input: [3, 4, 4, 3, 6, 3]

remove the 3 at index 0
remove the 4 at index 1
remove the 3 at index 3
Expected output: [4, 6, 3]

More examples can be found in the test cases.

Good luck!
*/
func SimpleRemoveDuplicates(arr []int) []int {
	repeated := make(map[int]int)
	result := make([]int, 0, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		if _, exists := repeated[arr[i]]; !exists {
			repeated[arr[i]] = len(result)
			result = append([]int{arr[i]}, result...)
		}
	}

	return result
}

func SimpleRemoveDuplicates1(a []int) (r []int) {
	m := map[int]int{}
	for _, v := range a {
		m[v]++
	}
	for _, v := range a {
		if m[v] == 1 {
			r = append(r, v)
		}
		m[v]--
	}
	return
}
