package kata8

/*
#Reversed sequence

Build a function that returns an array of integers from n to 1 where n>0.

Example : n=5 --> [5,4,3,2,1]
*/
func ReverseSeq(n int) []int {
	var result = make([]int, n)
	j := n
	for i := 1; i <= n; i++ {
		result[i-1] = j
		j--
	}
	return result
}

func ReverseSeq1(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i
	}
	return arr
}

func ReverseSeq2(n int) (result []int) {
	for i := n; i > 0; i-- {
		result = append(result, i)
	}
	return
}

func ReverseSeq3(n int) []int {
	ans := []int{}
	for n > 0 {
		ans = append(ans, n)
		n--
	}

	return ans
}
