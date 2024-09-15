package kata8

import "testing"

func TestLitres(t *testing.T) {
	testCases := []struct {
		time float64
		out  int
	}{
		{2, 1},
		{1.4, 0},
		{12.3, 6},
		{0.82, 0},
		{11.8, 5},
		{1787, 893},
		{0, 0},
	}

	for _, test := range testCases {
		got := Litres(test.time)
		want := test.out
		if want != got {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
