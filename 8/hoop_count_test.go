package kata8

import "testing"

func TestHoopCount(t *testing.T) {
	testCases := []struct {
		n   int
		out string
	}{
		{3, "Keep at it until you get it"},
		{12, "Great, now move on to tricks"},
	}

	for _, test := range testCases {
		got := HoopCount(test.n)
		want := test.out
		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
