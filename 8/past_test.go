package kata8

import "testing"

func TestPast(t *testing.T) {
	testCases := []struct {
		h        int
		m        int
		s        int
		expected int
	}{
		{0, 1, 1, 61000},
		{1, 1, 1, 3661000},
		{0, 0, 0, 0},
		{1, 0, 1, 3601000},
		{1, 0, 0, 3600000},
	}

	for _, test := range testCases {
		got := Past(test.h, test.m, test.s)
		if got != test.expected {
			t.Errorf("For input %v, %v, %v, expected %v but got %v", test.h, test.m, test.s, test.expected, got)
		}
		got = Past2(test.h, test.m, test.s)
		if got != test.expected {
			t.Errorf("For input %v, %v, %v, expected %v but got %v", test.h, test.m, test.s, test.expected, got)
		}
	}
}
