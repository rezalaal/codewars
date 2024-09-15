package kata7

import "testing"

func TestGps(t *testing.T) {
	testCases := []struct{
		s int
		x []float64
		out int
	}{
		{20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}, 41},
		{12, []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}, 219},
	}

	for _, test := range testCases {
		got := Gps(test.s, test.x)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}