package kata8

import (
	"testing"
)

func TestOddCount(t *testing.T) {
	testCases := []struct {
		n int
		out int
	}{
		{15,    7},
		{15023, 7511},
	}

	for _, test := range testCases {
		got := OddCount(test.n)
		expected := test.out

		if got != expected {
			t.Errorf("Expected %v but got %v", expected, got)
		}
	}
}