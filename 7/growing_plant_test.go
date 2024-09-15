package kata7

import "testing"

func TestGrowingPlant(t *testing.T) {
	testCases := []struct {
		upSpeed int
		downSpeed int
		desiredHeight int
		out int
	}{
		{100, 10, 910, 10},
		{10, 9, 4, 1},
		{5, 2, 5, 1},
		{5, 2, 6, 2},
	}

	for _, test := range testCases{
		got := GrowingPlant(test.upSpeed, test.downSpeed, test.desiredHeight)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = GrowingPlant1(test.upSpeed, test.downSpeed, test.desiredHeight)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
