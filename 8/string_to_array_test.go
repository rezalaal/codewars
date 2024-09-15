package kata8

import "testing"

func TestStringToArray(t *testing.T) {
	testCases := []struct{
		str string
		out []string
	}{
		{"Robin Singh",                        []string{"Robin", "Singh"}},
		{"CodeWars",                           []string{"CodeWars"}},
		{"I love arrays they are my favorite", []string{"I", "love", "arrays", "they", "are", "my", "favorite"}},
		{"1 2 3",                              []string{"1", "2", "3"}},
	}

	for _, test :=range testCases {
		got := StringToArray(test.str)
		expected := test.out

		for i, item :=range got {
			if expected[i] != item {
				t.Errorf("Expected %v but got %v", expected[i], item)
			}
		}
	}
}