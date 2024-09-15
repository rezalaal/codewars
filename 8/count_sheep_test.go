package kata8

import "testing"

func TestCountSheep(t *testing.T) {
	testCases := []struct {
		num int
		out string
	}{
		{2, "1 sheep...2 sheep..."},
		{0, ""},
		{1, "1 sheep..."},
	}

	for _, test := range testCases {
		got := CountSheep(test.num)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but want %v", want, got)
		}
	}
}
