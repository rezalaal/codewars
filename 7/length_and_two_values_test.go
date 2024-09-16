package kata7

import "testing"

func TestAlternate(t *testing.T) {
	testCases := []struct {
		n           int
		firstValue  string
		secondValue string
		out         []string
	}{
		{5, "true", "false", []string{"true", "false", "true", "false", "true"}},
		{20, "blue", "red", []string{"blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red", "blue", "red"}},
		{0, "", "", []string{}},
	}

	for _, test := range testCases {
		got := Alternate(test.n, test.firstValue, test.secondValue)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
