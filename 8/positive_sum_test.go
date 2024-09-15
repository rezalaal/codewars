package kata8

import "testing"

func TestPositiveSum(t *testing.T) {
	testCases := []struct {
		numbers []int
		out     int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, -2, 3, 4, 5}, 13},
		{[]int{}, 0},
		{[]int{-1, -2, -3, -4, -5}, 0},
		{[]int{-1, 2, 3, 4, -5}, 9},
	}

	for _, test := range testCases {
		got := PositiveSum(test.numbers)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
