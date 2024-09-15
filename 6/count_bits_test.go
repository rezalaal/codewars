package kata6

import "testing"

func TestCountBits(t *testing.T) {
	testCases := []struct {
		n   uint
		out int
	}{
		{0, 0},
		{4, 1},
		{7, 3},
		{9, 2},
		{10, 2},
	}

	for _, test := range testCases {
		got := CountBits(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = CountBits2(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = CountBits3(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = CountBits4(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
