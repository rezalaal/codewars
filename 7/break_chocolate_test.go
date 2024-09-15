package kata7

import "testing"

func TestBreakChocolate(t *testing.T) {
	testCases := []struct {
		n   int
		m   int
		out int
	}{
		{5, 5, 24},
		{1, 1, 0},
	}

	for _, test := range testCases {
		got := BreakChocolate(test.n, test.m)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = BreakChocolate1(test.n, test.m)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
