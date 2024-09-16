package kata7

import "testing"

func TestStantonMeasure(t *testing.T) {
	testCases := []struct {
		arr []int
		out int
	}{
		{[]int{1, 4, 3, 2, 1, 2, 3, 2}, 3},
		{[]int{1, 4, 3, 0, 1, 9, 3, 6}, 0},
	}

	for _, test := range testCases {
		got := StantonMeasure(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
