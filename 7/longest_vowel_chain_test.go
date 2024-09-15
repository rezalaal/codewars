package kata7

import "testing"

func TestLongestVowelChain(t *testing.T) {
	testCases := []struct {
		s   string
		out int
	}{
		{"codewarriors", 2},
		{"suoidea", 3},
		{"ultrarevolutionariees", 3},
		{"strengthlessnesses", 1},
		{"cuboideonavicuare", 2},
		{"chrononhotonthuooaos", 5},
		{"iiihoovaeaaaoougjyaw", 8},
	}

	for _, test := range testCases {
		got := LongestVowelChain(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
