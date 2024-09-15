package kata8

import "testing"

func TestHero(t *testing.T) {
	testCases := []struct {
		bullets int
		dragons int
		out     bool
	}{
		{10, 5, true},
		{7, 4, false},
		{4, 5, false},
		{100, 40, true},
		{1500, 751, false},
		{0, 1, false},
	}

	for _, test := range testCases {
		got := Hero(test.bullets, test.dragons)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = Hero2(test.bullets, test.dragons)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
