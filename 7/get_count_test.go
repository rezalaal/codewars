package kata7

import "testing"

func TestGetCount(t *testing.T) {
	testCases := []struct {
		str   string
		count int
	}{
		{"abracadabra", 5},
	}

	for _, test := range testCases {
		got := GetCount(test.str)
		want := test.count
		assertCorrectMessage(t, want, got)

		got = GetCount1(test.str)
		want = test.count
		assertCorrectMessage(t, want, got)

		got = GetCount2(test.str)
		want = test.count
		assertCorrectMessage(t, want, got)
	}
}
