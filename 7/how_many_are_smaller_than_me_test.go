package kata7

import "testing"

func TestSmaller(t *testing.T) {
	testCases := []struct {
		arr []int
		out []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{4, 3, 2, 1, 0}},
		{[]int{1, 2, 3}, []int{0, 0, 0}},
		{[]int{1, 2, 0}, []int{1, 1, 0}},
		{[]int{1, 2, 1}, []int{0, 1, 0}},
		{[]int{1, 1, -1, 0, 0}, []int{3, 3, 0, 0, 0}},
		{[]int{5, 4, 7, 9, 2, 4, 4, 5, 6}, []int{4, 1, 5, 5, 0, 0, 0, 0, 0}},
	}

	for _, test := range testCases {
		got := Smaller(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
