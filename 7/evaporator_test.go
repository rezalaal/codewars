package kata7

import "testing"

func TestEvaporator(t *testing.T) {
	testCases := []struct {
		content    float64
		evapPerDay int
		threshold  int
		out        int
	}{
		{10, 10, 10, 22},
		{10, 10, 5, 29},
		{100, 5, 5, 59},
	}

	for _, test := range testCases {
		got := Evaporator(test.content, test.evapPerDay, test.threshold)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Evaporator2(test.content, test.evapPerDay, test.threshold)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Evaporator3(test.content, test.evapPerDay, test.threshold)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
