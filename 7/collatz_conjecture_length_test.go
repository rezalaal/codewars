package kata7

import "testing"

func TestCollatz(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{20, 8},
		{15, 18},
		{100, 26},
		{10, 7},
		{500, 111},
		{1, 1},
		{1000000000, 101},
		{1000000000000000, 276},
	}

	for _, test := range testCases {
		got := Collatz(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
