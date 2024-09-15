package kata8

import "testing"

func TestCentury(t *testing.T) {
	testCases := []struct {
		year int
		output int
	}{
		{1990, 20},
		{1705, 18},
		{1900, 19},
		{1601, 17},
		{2000, 20},
		{89,   1},
	}

	for _, test :=range testCases {
		got := Century(test.year)
		expected := test.output

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Century2(test.year)
		expected = test.output

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = Century3(test.year)
		expected = test.output

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}