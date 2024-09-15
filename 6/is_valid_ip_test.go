package kata6

import "testing"

func TestIsValidIp(t *testing.T) {
	testCases := []struct {
		ip  string
		out bool
	}{
		{"12.255.56.1", true},
		{"", false},
		{"abc.def.ghi.jkl", false},
		{"123.456.789.0", false},
		{"12.34.56", false},
		{"12.34.56 .1", false},
		{"12.34.56.-1", false},
		{"123.045.067.089", false},
		{"127.1.1.0", true},
		{"0.0.0.0", true},
		{"0.34.82.53", true},
		{"192.168.1.300", false},
		{"h.v.l.s", false},
	}

	for _, test := range testCases {
		got := IsValidIp(test.ip)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
}
