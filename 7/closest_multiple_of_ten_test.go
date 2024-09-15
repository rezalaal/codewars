package kata7

import "testing"

func TestClosestMultipleOf10(t *testing.T) {
	testCases := []struct {
		n   uint32
		out uint32
	}{
		{22, 20},
		{25, 30},
		{37, 40},
		{54, 50},
		{55, 60},
	}

	for _, test := range testCases {
		got := ClosestMultipleOf10(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
