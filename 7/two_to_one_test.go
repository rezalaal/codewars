package kata7

import "testing"

func TestTwoToOne(t *testing.T) {
	testCases := []struct {
		s1  string
		s2  string
		out string
	}{
		{"aretheyhere", "yestheyarehere", "aehrsty"},
		{"loopingisfunbutdangerous", "lessdangerousthancoding", "abcdefghilnoprstu"},
	}

	for _, test := range testCases {
		got := TwoToOne(test.s1, test.s2)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = TwoToOne1(test.s1, test.s2)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = TwoToOne2(test.s1, test.s2)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
