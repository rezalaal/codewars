package kata7

import "testing"

func TestMxDifLg(t *testing.T) {
	testCases := []struct {
		a1  []string
		a2  []string
		out int
	}{
		{[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}, []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}, 13},
		{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, []string{"bbbaaayddqbbrrrv"}, 10},
	}

	for _, test := range testCases {
		got := MxDifLg(test.a1, test.a2)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = MxDifLg2(test.a1, test.a2)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = MxDifLg3(test.a1, test.a2)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
