package kata7

import "testing"

func TestHeron(t *testing.T) {
	testCases := []struct {
		a    float64
		b    float64
		c    float64
		area float64
	}{
		{3.0, 4.0, 5.0, 6.0},
	}

	for _, test := range testCases {
		got := Heron(test.a, test.b, test.c)
		want := test.area
		assertCorrectMessage(t, want, got)
	}
}
