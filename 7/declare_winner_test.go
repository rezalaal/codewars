package kata7

import "testing"

func TestDeclareWinner(t *testing.T) {
	testCases := []struct {
		fighter1      Fighter
		fighter2      Fighter
		firstAttacker string
		out           string
	}{
		{Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew", "Lew"},
		{Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry", "Harry"},
		{Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harry", "Harald"},
		{Fighter{"Harald", 20, 5}, Fighter{"Harry", 5, 4}, "Harald", "Harald"},
		{Fighter{"Jerry", 30, 3}, Fighter{"Harald", 20, 5}, "Jerry", "Harald"},
	}

	for _, test := range testCases {
		got := DeclareWinner(test.fighter1, test.fighter2, test.firstAttacker)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = DeclareWinner1(test.fighter1, test.fighter2, test.firstAttacker)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = DeclareWinner2(test.fighter1, test.fighter2, test.firstAttacker)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
