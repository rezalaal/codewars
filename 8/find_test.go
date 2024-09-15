package kata8

import "testing"

func TestFind(t *testing.T) {
	array := []int{2, 3, 5, 7, 11}
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"Find 5", 5, "2"},
		{"Find 11", 11, "4"},
		{"Find 3", 3, "1"},
		{"Find 2", 2, "0"},
		{"Find 7", 7, "3"},
		{"Find 1", 1, "Not found"},
		{"Find 8", 8, "Not found"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Find(array, test.input)
			if got != test.expected {
				t.Errorf("Find(%d): expected %v but got %v", test.input, test.expected, got)
			}
		})
	}
}
