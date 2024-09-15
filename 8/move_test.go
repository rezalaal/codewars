package kata8

import (
	"testing"
)

func TestMove(t *testing.T) {
	testCases := []struct {
		position int
		roll     int
		out      int
	}{
		{0, 4, 8},
		{3, 6, 15},
		{2, 5, 12},
	}

	for _, test := range testCases {
		got := Move(test.position, test.roll)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but want %v", want, got)
		}
	}
}
