package kata8

import "testing"

func TestTwoSort(t *testing.T) {
	testCases := []struct {
		arr []string
		out string
	}{
		{[]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}, "b***i***t***c***o***i***n"},
		{[]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}, "a***r***e"},
		{[]string{"lets", "talk", "about", "javascript", "the", "best", "language"}, "a***b***o***u***t"},
		{[]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}, "c***o***d***e"},
		{[]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}, "L***e***t***s"},
	}

	for _, test :=range testCases {
		got := TwoSort(test.arr)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
