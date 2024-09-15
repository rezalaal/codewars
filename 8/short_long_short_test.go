package kata8

import "testing"

func TestShortLongShort(t *testing.T) {
	testCases := []struct {
		a   string
		b   string
		out string
	}{
		{"45", "1", "1451"},
		{"13", "200", "1320013"},
		{"Soon", "Me", "MeSoonMe"},
		{"U", "False", "UFalseU"},
	}

	for _, test := range testCases {
		got := ShortLongShort(test.a, test.b)
		want := test.out

		if got != want {
			t.Errorf("Expected %v but got %v", want, got)
		}
	}
}
