package kata7

import "testing"

func TestWordPattern(t *testing.T) {
	testCases := []struct{
		word string
		out string
	}{
		{"hello", "0.1.2.2.3"},
        {"heLlo", "0.1.2.2.3"},
        {"helLo", "0.1.2.2.3"},
        {"Hippopotomonstrosesquippedaliophobia", "0.1.2.2.3.2.3.4.3.5.3.6.7.4.8.3.7.9.7.10.11.1.2.2.9.12.13.14.1.3.2.0.3.15.1.13"},
	}

	for _, test := range testCases {
		got := WordPattern(test.word)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}