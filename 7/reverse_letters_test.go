package kata7

import "testing"

func TestReverseLetters(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"krishan", "nahsirk"},
		{"ultr53o?n", "nortlu"},
		{"ab23c", "cba"},
		{"krish21an", "nahsirk"},
		{"", ""},
		{"Az", "zA"},
	}

	for _, test := range testCases {
		got := ReverseLetters(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = ReverseLetters1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = ReverseLetters2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = ReverseLetters3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
