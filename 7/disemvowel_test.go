package kata7

import "testing"

func TestDisemvowel(t *testing.T) {
	testCases := []struct {
		comment string
		out     string
	}{
		{"This website is for losers LOL!", "Ths wbst s fr lsrs LL!"},
	}

	for _, test := range testCases {
		got := Disemvowel(test.comment)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Disemvowel1(test.comment)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Disemvowel2(test.comment)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Disemvowel3(test.comment)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Disemvowel4(test.comment)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
