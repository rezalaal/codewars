package kata7

import "testing"

func TestMaxMultiple(t *testing.T) {
	testCases := []struct {
		d   int
		b   int
		out int
	}{
		{2, 7, 6},
		{3, 10, 9},
		{7, 17, 14},
		{10, 50, 50},
		{37, 200, 185},
		{7, 100, 98},
	}

	for _, test := range testCases {
		got := MaxMultiple(test.d, test.b)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = MaxMultiple2(test.d, test.b)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = MaxMultiple3(test.d, test.b)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = MaxMultiple4(test.d, test.b)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
