package kata8

import "testing"

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		srt string
		out bool
	}{
		{"a", true},
		{"aba", true},
		{"Abba", true},
		{"hello", false},
	}

	for _, test := range testCases {
		got := IsPalindrome(test.srt)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
