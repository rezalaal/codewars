package kata8

import "testing"

func TestIsDivisible(t *testing.T) {
	testCases := []struct {
		n   int
		x   int
		y   int
		out bool
	}{
		{3, 3, 4, false},
		{12, 3, 4, true},
		{8, 3, 4, false},
		{48, 3, 4, true},
		{100, 5, 10, true},
		{100, 5, 3, false},
		{4, 4, 2, true},
		{5, 2, 3, false},
		{17, 17, 17, true},
		{17, 1, 17, true},
	}

	for _, test := range testCases {
		got := IsDivisible(test.n, test.x, test.y)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
