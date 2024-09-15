package kata7

import "testing"

func TestContainAllRots(t *testing.T) {
	testCases := []struct {
		strng string
		arr   []string
		out   bool
	}{
		{"bsjq", []string{"bsjq", "qbsj", "sjqb", "twZNsslC", "jqbs"}, true},
		{"XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}, false},
		{"", []string{}, true},
	}

	for _, test := range testCases {
		got := ContainAllRots(test.strng, test.arr)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
