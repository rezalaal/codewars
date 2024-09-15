package kata7

import "testing"

func TestSpacify(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"hello world", "h e l l o   w o r l d"},
		{"12345", "1 2 3 4 5"},
		{"Pippi", "P i p p i"},
		{"a", "a"},
		{"", ""},
	}

	for _, test := range testCases {
		got := Spacify(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Spacify1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
