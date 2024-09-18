package kata7

import "testing"

func TestConsecutiveLetters(t *testing.T) {
	testCases := []struct {
		s   string
		out bool
	}{
		{"abc", true},
		{"abd", false},
		{"dabc", true},
		{"abbc", false},
	}

	for _, test := range testCases {
		got := ConsecutiveLetters(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
