package kata8

import "testing"

func TestRps(t *testing.T) {
	testCases := []struct {
		p1  string
		p2  string
		out string
	}{
		{"rock", "scissors", "Player 1 won!"},
		{"scissors", "rock", "Player 2 won!"},
		{"rock", "rock", "Draw!"},
	}

	for _, test := range testCases {
		got := Rps(test.p1, test.p2)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
