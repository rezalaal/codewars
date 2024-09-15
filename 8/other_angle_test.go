package kata8

import "testing"

func TestOtherAngle(t *testing.T) {
	testCases := []struct {
		a   int
		b   int
		out int
	}{
		{30, 60, 90},
		{60, 60, 60},
		{43, 78, 59},
		{10, 20, 150},
	}

	for _, test := range testCases {
		got := OtherAngle(test.a, test.b)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}