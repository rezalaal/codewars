package kata7

import "testing"

func TestRowSumOddNumbers(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{2, 8},
		{13, 2197},
		{19, 6859},
		{41, 68921},
		{42, 74088},
		{74, 405224},
		{86, 636056},
		{93, 804357},
		{101, 1030301},
	}

	for _, test := range testCases {
		got := RowSumOddNumbers(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = RowSumOddNumbers1(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = RowSumOddNumbers2(test.n)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
