package kata7

import "testing"

func TestHighAndLow(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"1 2 3", "3 1"},
	}

	for _, test := range testCases {
		got := HighAndLow(test.in)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = HighAndLow2(test.in)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
