package kata7

import "testing"

func TestScale(t *testing.T) {
	testCases := []struct {
		s   string
		k   int
		n   int
		out string
	}{
		{"abcd\nefgh\nijkl\nmnop", 2, 3, "aabbccdd\naabbccdd\naabbccdd\neeffgghh\neeffgghh\neeffgghh\niijjkkll\niijjkkll\niijjkkll\nmmnnoopp\nmmnnoopp\nmmnnoopp"},
		{"", 5, 5, ""},
		{"Kj\nSH", 1, 2, "Kj\nKj\nSH\nSH"},
	}

	for _, test := range testCases {
		got := Scale(test.s, test.k, test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
