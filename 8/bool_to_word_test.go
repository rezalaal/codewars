package kata8

import "testing"

func TestBoolToWord(t *testing.T) {
	testCases := []struct {
		word bool
		out string
	}{
		{true,  "Yes"},
		{false, "No"},
	}

	for _, test :=range testCases {
		got := BoolToWord(test.word)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = BoolToWord2(test.word)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}