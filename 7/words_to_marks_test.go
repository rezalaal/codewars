package kata7

import "testing"

func TestWordsToMarks(t *testing.T) {
	testCases := []struct {
		s   string
		out int
	}{
		{"attitude", 100},
		{"friends", 75},
		{"family", 66},
		{"selfness", 99},
		{"knowledge", 96},
	}

	for _, test := range testCases {
		got := WordsToMarks(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = WordsToMarks2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = WordsToMarks3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = WordsToMarks4(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = WordsToMarks5(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
