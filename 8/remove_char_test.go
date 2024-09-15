package kata8

import "testing"

func TestRemoveChar(t *testing.T) {
	testCases := []struct {
		word string
		out  string
	}{
		{"eloquent", "loquen"},
		{"country", "ountr"},
		{"person", "erso"},
		{"place", "lac"},
		{"", ""},
		{"a", ""},
	}
	for _, test := range testCases {
		got := RemoveChar(test.word)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = RemoveChar2(test.word)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = RemoveChar3(test.word)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = RemoveChar4(test.word)
		want = test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
