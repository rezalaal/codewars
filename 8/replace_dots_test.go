package kata8

import "testing"

func TestReplaceDots(t *testing.T) {
	testCases := []struct {
		str string
		out string
	}{
		{"", ""},
		{"no dots", "no dots"},
		{"one.two.three", "one-two-three"},
		{"...", "---"},
	}

	for _, test := range testCases {
		got := ReplaceDots(test.str)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = ReplaceDots1(test.str)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = ReplaceDots2(test.str)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
