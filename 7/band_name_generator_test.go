package kata7

import "testing"

func TestBandNameGenerator(t *testing.T) {
	testCases := []struct {
		word string
		out  string
	}{
		{"knife", "The Knife"},
		{"tart", "Tartart"},
		{"sandles", "Sandlesandles"},
		{"bed", "The Bed"},
	}

	for _, test := range testCases {
		got := bandNameGenerator(test.word)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
