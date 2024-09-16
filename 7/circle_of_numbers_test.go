package kata7

import "testing"

func TestCircleOfNumbers(t *testing.T) {
	testCases := []struct {
		n           int
		firstNumber int
		out         int
	}{
		{10, 2, 7},
		{10, 7, 2},
		{4, 1, 3},
		{6, 3, 0},
		{20, 0, 10},
	}

	for _, test := range testCases {
		got := CircleOfNumbers(test.n, test.firstNumber)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
