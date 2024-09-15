package kata8

import "testing"

func TestNthEven(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 0},
		{2, 2},
		{3, 4},
		{100, 198},
		{1298734, 2597466},
	}

	for _, test := range testCases {
		got := NthEven(test.n)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
