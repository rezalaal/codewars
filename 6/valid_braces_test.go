package kata6

import "testing"

func TestValidBraces(t *testing.T) {
	testCases := []struct {
		str string
		out bool
	}{
		{"(){}[]", true},
		{"([{}])", true},
		{"(}", false},
		{"[(])", false},
		{"[({)](]", false},
	}

	for _, test := range testCases {
		got := ValidBraces(test.str)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = ValidBraces1(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = ValidBraces2(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = ValidBraces3(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
