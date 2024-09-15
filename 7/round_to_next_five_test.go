package kata7

import "testing"

func TestRoundToNextFive(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{0, 0},
		{1, 5},
		{-1, 0},
		{5, 5},
		{7, 10},
		{20, 20},
		{39, 40},
		{121, 125},
		{555, 555},
		{-97932, -97930},
	}

	for _, test := range testCases {
		got := RoundToNextFive(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = RoundToNextFive1(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = RoundToNextFive2(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = RoundToNextFive3(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
