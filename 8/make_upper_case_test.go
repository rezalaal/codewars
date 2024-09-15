package kata8

import "testing"

func TestMakeUpperCase(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"hello", "HELLO"},
		{"hello world", "HELLO WORLD"},
		{"hello world !", "HELLO WORLD !"},
		{"heLlO wORLd !", "HELLO WORLD !"},
		{"1,2,3 hello world!", "1,2,3 HELLO WORLD!"},
	}

	for _, test := range testCases {
		got := MakeUpperCase(test.s)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
