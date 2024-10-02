package kata7

/*
Reverse it
Write reverseList function that simply reverses lists.
*/
func ReverseList(lst []int) []int {
	res := []int{}
	for i := len(lst) - 1; i >= 0; i-- {
		res = append(res, lst[i])
	}
	return res
}

func ReverseList1(lst []int) []int {
	for i := 0; i < len(lst)/2; i++ {
		lst[i], lst[len(lst)-1-i] = lst[len(lst)-1-i], lst[i]
	}
	return lst
}

func ReverseList2(lst []int) []int {
	head, tail := 0, len(lst)-1
	for head < tail {
		lst[head], lst[tail] = lst[tail], lst[head]
		head++
		tail--
	}
	return lst
}

func ReverseList3(lst []int) []int {

	outputList := []int{}
	for _, char := range lst {
		outputList = append([]int{char}, outputList...)
	}

	return outputList
}
