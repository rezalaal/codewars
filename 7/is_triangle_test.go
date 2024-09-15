package kata7

import "testing"

func TestIsTriangle(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		c   int
		out bool
	}{
		{5, 1, 2, false},
		{1, 2, 5, false},
		{2, 5, 1, false},
		{4, 2, 3, true},
		{5, 1, 5, true},
		{2, 2, 2, true},
		{-1, 2, 3, false},
	}

	for _, test := range testCases {
		got := IsTriangle(test.a, test.b, test.c)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = IsTriangle1(test.a, test.b, test.c)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
