package kata8

import "testing"

func TestReverseWords(t *testing.T) {
	testCases := []struct {
		str string
		out string
	}{
		{"hello world!", "world! hello"},
		{"yoda doesn't speak like this", "this like speak doesn't yoda"},
		{"foobar", "foobar"},
		{"kata editor", "editor kata"},
		{"row row row your boat", "boat your row row row"},
	}

	for _, test := range testCases {
		got := ReverseWords(test.str)
		want := test.out

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	}
}
