package kata8

import "testing"

func TestSquareSum(t *testing.T) {
	testCases := []struct {
		numbers []int
		out     int
	}{
		{[]int{1, 2}, 5},
		{[]int{0, 3, 4, 5}, 50},
		{[]int{}, 0},
	}
	for _, test := range testCases {
		got := SquareSum(test.numbers)
		want := test.out

		if got != want {
			t.Errorf("Expect %v but got %v", want, got)
		}
	}
}
