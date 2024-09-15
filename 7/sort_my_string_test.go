package kata7

import "testing"

func TestSortMyString(t *testing.T) {
	testCases := []struct{
		s string
		out string
	}{
		{"CodeWars",      "CdWr oeas"},
		{"YCOLUE'VREER",  "YOU'RE CLEVER"},
	}

	for _, test := range testCases {
		got := SortMyString(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}