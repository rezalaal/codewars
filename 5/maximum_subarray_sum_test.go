package kata5

import "testing"

func TestMaximumSubarraySum(t *testing.T) {
	testCases := []struct {
		numbers []int
		out     int
	}{
		{[]int{}, 0},
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}, 0},
	}

	for _, test := range testCases {
		got := MaximumSubarraySum(test.numbers)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
