package kata6

/*
#Array.diff

Your goal in this kata is to implement a difference function, which subtracts one list from another and returns the result.

It should remove all values from list a, which are present in list b keeping their order.

array_diff({1, 2}, 2, {1}, 1, *z) == {2} (z == 1)
If a value is present in b, all of its occurrences must be removed from the other:

array_diff({1, 2, 2, 2, 3}, 5, {2}, 1, *z) == {1, 3} (z == 2)
*/
func ArrayDiff1(a, b []int) []int {
	result := []int{}
	if len(b) == 0 {
		return a
	}
	for _, item := range a {
		inB := false
		for _, member := range b {
			if item == member {
				inB = true
			}

		}
		if inB == false {
			result = append(result, item)
		}
	}
	return result
}

func ArrayDiff(a, b []int) (r []int) {
	m := map[int]bool{}
	for _, v := range b {
		m[v] = true
	}
	for _, v := range a {
		if !m[v] {
			r = append(r, v)
		}
	}
	return // a-b
}
