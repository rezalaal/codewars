package kata7

import "testing"

func TestAddLetters(t *testing.T) {
	testCases := []struct {
		letters []rune
		out     rune
	}{
		{[]rune{'a', 'b', 'c'}, 'f'},
		{[]rune{'z'}, 'z'},
		{[]rune{'a', 'b'}, 'c'},
		{[]rune{'c'}, 'c'},
		{[]rune{'z', 'a'}, 'a'},
		{[]rune{'y', 'c', 'b'}, 'd'},
		{[]rune{}, 'z'},
	}

	for _, test := range testCases {
		got := AddLetters(test.letters)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}
