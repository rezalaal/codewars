package kata7

import "testing"

func TestFib(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
	}

	for _, test := range testCases {
		got := Fib(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
