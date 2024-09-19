package kata7

import "testing"

func TestBubblesortOnce(t *testing.T) {
	testCases := []struct {
		numbers []int
		out     []int
	}{
		{[]int{9, 7, 5, 3, 1, 2, 4, 6, 8}, []int{7, 5, 3, 1, 2, 4, 6, 8, 9}},
	}

	for _, test := range testCases {
		got := BubblesortOnce(test.numbers)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
