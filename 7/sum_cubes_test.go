package kata7

import "testing"

func TestSumCubes(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{2, 9},
		{3, 36},
		{4, 100},
		{10, 3025},
		{123, 58155876},
	}

	for _, test := range testCases {
		got := SumCubes(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SumCubes2(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
