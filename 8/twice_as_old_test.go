package kata8

import "testing"

func TestTwiceAsOld(t *testing.T) {
	testCases := []struct {
		dadYearsOld int
		sonYearsOld int
		out         int
	}{
		{36, 7, 22},
		{55, 30, 5},
		{42, 21, 0},
		{22, 1, 20},
		{29, 0, 29},
	}

	for _, test := range testCases {
		got := TwiceAsOld(test.dadYearsOld, test.sonYearsOld)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
