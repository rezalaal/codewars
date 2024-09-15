package kata7

import "testing"

func TestIndexedCapitalization(t *testing.T) {
	testCases := []struct {
		st  string
		arr []int
		out string
	}{
		{"abcdef", []int{1, 2, 5}, "aBCdeF"},
		{"abcdef", []int{1, 2, 5, 100}, "aBCdeF"},
		{"codewars", []int{1, 3, 5, 50}, "cOdEwArs"},
		{"abracadabra", []int{2, 6, 9, 10}, "abRacaDabRA"},
		{"codewarriors", []int{5}, "codewArriors"},
		{"indexinglessons", []int{0}, "Indexinglessons"},
	}

	for _, test := range testCases {
		got := IndexedCapitalization(test.st, test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
