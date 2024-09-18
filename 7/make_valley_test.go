package kata7

import "testing"

func TestMakeValley(t *testing.T) {
	testCases := []struct {
		arr []int
		out []int
	}{
		{[]int{17, 17, 15, 14, 8, 7, 7, 5, 4, 4, 1}, []int{17, 15, 8, 7, 4, 1, 4, 5, 7, 14, 17}},
		{[]int{20, 7, 6, 2}, []int{20, 6, 2, 7}},
		{[]int{14, 10, 8}, []int{14, 8, 10}},
		{[]int{20, 18, 17, 13, 12, 12, 10, 9, 4, 2, 2, 1, 1}, []int{20, 17, 12, 10, 4, 2, 1, 1, 2, 9, 12, 13, 18}},
	}

	for _, test := range testCases {
		got := MakeValley(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
