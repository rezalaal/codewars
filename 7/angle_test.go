package kata7

import "testing"

func TestAngle(t *testing.T) {
	testCases := []struct {
		n   int
		out int
	}{
		{3, 180},
		{4, 360},
	}

	for _, test := range testCases {
		got := Angle(test.n)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
