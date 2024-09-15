package kata7

import "testing"

func TestPrinterError(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"aaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "3/56"},
		{"kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyz", "6/60"},
		{"kkkwwwaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbmmmmmmmmmmmmmmmmmmmxyzuuuuu", "11/65"},
	}

	for _, test := range testCases {
		got := PrinterError(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = PrinterError1(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = PrinterError2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = PrinterError3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
