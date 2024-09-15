package kata8

import "testing"

func TestAmIWilson(t *testing.T) {
	testCases := []struct {
		n   int
		out bool
	}{
		{0, false},
		{1, false},
		{5, true},
		{8, false},
		{9, false},
	}

	for _, test := range testCases {
		got := AmIWilson(test.n)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
