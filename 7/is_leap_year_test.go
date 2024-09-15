package kata7

import "testing"

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		year int
		out  bool
	}{
		{2020, true},
		{2000, true},
		{2015, false},
		{2100, false},
	}

	for _, test := range testCases {
		got := IsLeapYear(test.year)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = IsLeapYear1(test.year)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
