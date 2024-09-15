package kata7

import "testing"

func TestDigits(t *testing.T) {
	testCases := []struct {
		n   uint64
		out int
	}{
		{5, 1},
		{12345, 5},
		{9876543210, 10},
	}

	for _, test := range testCases {
		got := Digits(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Digits1(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Digits2(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Digits3(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
