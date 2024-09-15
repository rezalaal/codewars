package kata8

import "testing"

func TestNearestSq(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{1, 1},
		{2, 1},
		{10, 9},
		{111, 121},
		{9999, 10000},
	}

	for _, test := range testCases {
		got := NearestSq(test.n)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}

		got = NearestSq1(test.n)
		expected = test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}
