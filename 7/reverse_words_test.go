package kata7

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		str string
		out string
	}{
		{"The quick brown fox jumps over the lazy dog.", "ehT kciuq nworb xof spmuj revo eht yzal .god"},
		{"apple", "elppa"},
		{"a b c d", "a b c d"},
		{"double  spaced  words", "elbuod  decaps  sdrow"},
		{"stressed desserts", "desserts stressed"},
		{"hello hello", "olleh olleh"},
	}

	for _, test := range testCases {
		got := ReverseWords(test.str)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = ReverseWords2(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
