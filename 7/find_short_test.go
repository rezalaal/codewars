package kata7

import "testing"

func TestFindShort(t *testing.T) {
	testCases := []struct {
		s   string
		out int
	}{
		{"bitcoin take over the world maybe who knows perhaps", 3},
		{"turns out random test cases are easier than writing out basic ones", 3},
		{"Let's travel abroad shall we", 2},
	}

	for _, test := range testCases {
		got := FindShort(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = FindShort1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = FindShort2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
