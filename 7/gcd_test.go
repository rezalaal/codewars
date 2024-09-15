package kata7

import "testing"

func TestGcd(t *testing.T) {
	testCases := []struct{
		x uint32
		y uint32
		out uint32
	}{
		{30, 12, 6},
		{8, 9, 1},
		{1, 1, 1},
	}

	for _, test := range testCases {
		got := Gcd(test.x, test.y)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}