package kata8

import "testing"

func TestCheckForFactor(t *testing.T) {
	testCases := []struct {
		base   int
		factor int
		out    bool
	}{
		{10, 2, true},
		{63, 7, true},
		{2450, 5, true},
		{24612, 3, true},
		{9, 2, false},
		{653, 7, false},
		{2453, 5, false},
		{24617, 3, false},
	}

	for _, test := range testCases {
		got := CheckForFactor(test.base, test.factor)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
