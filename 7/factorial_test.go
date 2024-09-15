package kata7

import "testing"

func TestFactorial(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{0, 1},
		{1, 1},
		{4, 24},
		{7, 5040},
		{17, 355687428096000},
	}

	for _, test := range testCases {
		got := Factorial(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Factorial1(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
