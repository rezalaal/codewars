package kata8

import "testing"

func TestExpressionMatter1(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		c   int
		out int
	}{
		{2, 1, 2, 6},
		{2, 1, 1, 4},
		{1, 1, 1, 3},
		{1, 2, 3, 9},
		{1, 3, 1, 5},
		{2, 2, 2, 8},
		{5, 1, 3, 20},
		{3, 5, 7, 105},
		{5, 6, 1, 35},
		{1, 6, 1, 8},
		{2, 6, 1, 14},
		{6, 7, 1, 48},
		{2, 10, 3, 60},
		{1, 8, 3, 27},
		{9, 7, 2, 126},
		{1, 1, 10, 20},
		{9, 1, 1, 18},
		{10, 5, 6, 300},
	}

	for _, test := range testCases {
		got := ExpressionMatter(test.a, test.b, test.c)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
