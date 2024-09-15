package kata7

import "testing"

func TestSortByLength(t *testing.T) {
	testCases := []struct {
		arr []string
		out []string
	}{
		{[]string{"beg", "life", "i", "to"}, []string{"i", "to", "beg", "life"}},
		{[]string{"", "moderately", "brains", "pizza"}, []string{"", "pizza", "brains", "moderately"}},
		{[]string{"longer", "longest", "short"}, []string{"short", "longer", "longest"}},
		{[]string{"dog", "food", "a", "of"}, []string{"a", "of", "dog", "food"}},
		{[]string{"", "dictionary", "eloquent", "bees"}, []string{"", "bees", "eloquent", "dictionary"}},
		{[]string{"a longer sentence", "the longest sentence", "a short sentence"}, []string{"a short sentence", "a longer sentence", "the longest sentence"}},
	}

	for _, test := range testCases {
		got := SortByLength(test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = SortByLength1(test.arr)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
