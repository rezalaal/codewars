package kata7

import (
	"testing"
)

// TestOrderedCount tests the OrderedCount function.
func TestOrderedCount(t *testing.T) {
	tests := []struct {
		input    string
		expected []Tuple
	}{
		{"abracadabra", []Tuple{{'a', 5}, {'b', 2}, {'r', 2}, {'c', 1}, {'d', 1}}},
		{"Code Wars",   []Tuple{{'C', 1}, {'o', 1}, {'d', 1}, {'e', 1}, {' ', 1}, {'W', 1}, {'a', 1}, {'r', 1}, {'s', 1}}},
		{"",            []Tuple{}},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := OrderedCount(test.input)
			if !equal(result, test.expected) {
				t.Errorf("OrderedCount(%q) = %v; want %v", test.input, result, test.expected)
			}
		})
	}
}

// equal checks if two slices of Tuple are equal.
func equal(a, b []Tuple) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}