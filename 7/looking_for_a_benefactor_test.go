package kata7

import "testing"

func TestNewAvg(t *testing.T) {
	testCases := []struct {
		arr  []float64
		navg float64
		out  int64
	}{
		{[]float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 16.0}, 90, 628},
		{[]float64{14, 30, 5, 7, 9, 11, 15}, 92, 645},
		{[]float64{1400.25, 30000.76, 5.56, 7, 9, 11, 15.48, 120.98}, 2000, -1},
	}

	for _, test := range testCases {
		got := NewAvg(test.arr, test.navg)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
