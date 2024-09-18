package kata7

import "testing"

func TestDominator(t *testing.T) {
	testCases := []struct {
		a   []int
		out int
	}{
		{[]int{3, 4, 3, 2, 3, 1, 3, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, -1},
		{[]int{1, 1, 1, 2, 2, 2}, -1},
		{[]int{1, 1, 1, 2, 2, 2, 2}, 2},
	}

	for _, test := range testCases {
		got := Dominator(test.a)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
