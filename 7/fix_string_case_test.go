package kata7

import "testing"

func TestFixStringCase(t *testing.T) {
	testCases := []struct {
		str string
		out string
	}{
		{"a", "a"},
		{"Z", "Z"},
		{"coDe", "code"},
		{"CODe", "CODE"},
		{"coDE", "code"},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
	}

	for _, test := range testCases {
		got := FixStringCase(test.str)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
