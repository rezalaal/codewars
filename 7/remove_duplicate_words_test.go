package kata7

import "testing"

func TestRemoveDuplicateWords(t *testing.T) {
	testCases := []struct {
		str string
		out string
	}{
		{"alpha beta beta gamma gamma gamma delta alpha beta beta gamma gamma gamma delta", "alpha beta gamma delta"},
		{"my cat is my cat fat", "my cat is fat"},
	}

	for _, test := range testCases {
		got := RemoveDuplicateWords(test.str)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = RemoveDuplicateWords1(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = RemoveDuplicateWords2(test.str)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
