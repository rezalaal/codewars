package kata7

import "testing"

func TestSimpleStringReversal(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"codewars", "srawedoc"},
		{"your code", "edoc ruoy"},
		{"your code rocks", "skco redo cruoy"},
	}

	for _, test := range testCases {
		got := SimpleStringReversal(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
