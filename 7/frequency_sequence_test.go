package kata7

import "testing"

func TestFreqSeq(t *testing.T) {
	testCases := []struct {
		str string
		sep string
		out string
	}{
		{"hello world", "-", "1-1-3-3-2-1-1-2-1-3-1"},
		{"19999999", ":", "1:7:7:7:7:7:7:7"},
		{"^^^**$", "x", "3x3x3x2x2x1"},
	}

	for _, test := range testCases {
		got := FreqSeq(test.str, test.sep)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
