package kata7

import "testing"

func TestStrong(t *testing.T) {
	testCases := []struct {
		n   int
		out string
	}{
		{1, "STRONG!!!!"},
		{2, "STRONG!!!!"},
		{145, "STRONG!!!!"},
		{7, "Not Strong !!"},
		{93, "Not Strong !!"},
		{185, "Not Strong !!"},
	}

	for _, test := range testCases {
		got := Strong(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
