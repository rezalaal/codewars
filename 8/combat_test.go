package kata8

import "testing"

func TestCombat(t *testing.T) {
	testCases := []struct {
		health float64
		damage float64
		out    float64
	}{
		{100.0, 50.0, 50.0},
		{1.0, 100, 0.0},
	}

	for _, test := range testCases {
		got := Combat(test.health, test.damage)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}

		got = Combat1(test.health, test.damage)
		want = test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
