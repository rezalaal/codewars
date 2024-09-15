package kata7

import "testing"

func TestInAscOrder(t *testing.T) {		
	testCases := []struct{
		numbers []int
		out bool
	}{
		{[]int{1, 2, 4, 7, 19},            true},
		{[]int{1, 2, 3, 4, 5},             true},
		{[]int{1, 6, 10, 18, 2, 4, 20},    false},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, false},
	}

	for _, test := range testCases {
		got := InAscOrder(test.numbers)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}