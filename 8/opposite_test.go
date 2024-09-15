package kata8

import "testing"

func TestOpposite(t *testing.T) {
	testCases := []struct {
		value int
		out   int
	}{
		{1, -1},
		{-1, 1},
		{0, 0},
	}

	for _, test := range testCases {
		got := Opposite(test.value)
		want := test.out
		if got != want {
			t.Errorf("Expect %v but got %v", want, got)
		}
	}
}
