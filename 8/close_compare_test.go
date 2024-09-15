package kata8

import "testing"

func TestCloseCompare(t *testing.T) {
	testCases := []struct {
		a float64
		b float64
		margin float64
		output int
	}{
		{4.0, 5.0, 0.0, -1},
		{5.0, 5.0, 0.0, 0},
		{6.0, 5.0, 0.0, 1},
		{2.0, 5.0, 3.0, 0},
		{5.0, 5.0, 3.0, 0},
		{8.0, 5.0, 3.0, 0},
		{8.1, 5.0, 3.0, 1},
		{1.99, 5.0, 3.0, -1},
	}

	for _, test :=range testCases {
		got := CloseCompare(test.a, test.b, test.margin)
		expected := test.output
		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}