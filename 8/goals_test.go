package kata8

import "testing"

func TestGoals(t *testing.T) {
	testCases := []struct {
		laLigaGoals          int
		copaDelReyGoals      int
		championsLeagueGoals int
		out                  int
	}{
		{0, 0, 0, 0},
		{43, 10, 5, 58},
	}

	for _, test := range testCases {
		got := Goals(test.laLigaGoals, test.copaDelReyGoals, test.championsLeagueGoals)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
