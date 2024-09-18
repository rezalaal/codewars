package kata7

import "testing"

func TestModifyMultiply(t *testing.T) {
	testCases := []struct {
		str string
		loc int
		num int
		out string
	}{
		{"This is a string", 3, 5, "string-string-string-string-string"},
		{"Creativity is the process of having original ideas that have value. It is a process; it's not random.", 8, 10, "that-that-that-that-that-that-that-that-that-that"},
		{"Self-control means wanting to be effective at some random point in the infinite radiations of my spiritual existence", 1, 1, "means"},
		{"Is sloppiness in code caused by ignorance or apathy? I don't know and I don't care.", 6, 8, "ignorance-ignorance-ignorance-ignorance-ignorance-ignorance-ignorance-ignorance"},
		{"Everything happening around me is very random. I am enjoying the phase, as the journey is far more enjoyable than the destination.", 2, 5, "around-around-around-around-around"},
		{"Go is fun", 0, 2,"Go-Go"},
		{"Test", 2, 0, ""},
	}

	for _, test := range testCases {
		got := ModifyMultiply(test.str, test.loc, test.num)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
