package kata7

import "testing"

func TestSequenceSum(t *testing.T) {
	testCases := []struct {
		start int
		end   int
		step  int
		out   int
	}{
		{2, 6, 2, 12},
		{1, 5, 1, 15},
		{1, 5, 3, 5},
		{0, 15, 3, 45},
		{16, 15, 3, 0},
		{2, 24, 22, 26},
		{2, 2, 2, 2},
		{2, 2, 1, 2},
		{1, 15, 3, 35},
		{15, 1, 3, 0},
	}

	for _, test := range testCases {
		got := SequenceSum(test.start, test.end, test.step)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SequenceSum1(test.start, test.end, test.step)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
