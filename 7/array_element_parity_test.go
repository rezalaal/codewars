package kata7

import "testing"

func TestArrayElementParity(t *testing.T) {
	testCases := []struct {
		arr []int
		out int
	}{
		{[]int{1, -1, 2, -2, 3}, 3},
		{[]int{-3, 1, 2, 3, -1, -4, -2}, -4},
		{[]int{1, -1, 2, -2, 3, 3}, 3},
		{[]int{-110, 110, -38, -38, -62, 62, -38, -38, -38}, -38},
		{[]int{-9, -105, -9, -9, -9, -9, 105}, -9},
	}

	for _, test := range testCases {
		got := ArrayElementParity(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
