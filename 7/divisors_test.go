package kata7

import "testing"

func TestDivisors(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{10, 4},
		{11, 2},
		{54, 8},
		{64, 7},
	}

	for _, test := range testCases {
		got := Divisors(test.n)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}
