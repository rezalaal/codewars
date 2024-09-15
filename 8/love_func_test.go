package kata8

import "testing"

func TestLoveFunc(t *testing.T) {
	testCases := []struct {
		flower1 int
		flower2 int
		out     bool
	}{
		{1, 4, true},
		{2, 2, false},
		{0, 1, true},
		{0, 0, false},
		{95, 721, false},
	}
	for _, test := range testCases {
		got := LoveFunc2(test.flower1, test.flower2)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
