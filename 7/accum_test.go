package kata7

import "testing"

func TestAccum(t *testing.T) {
	testCases := []struct {
		s   string
		out string
	}{
		{"ZpglnRxqenU", "Z-Pp-Ggg-Llll-Nnnnn-Rrrrrr-Xxxxxxx-Qqqqqqqq-Eeeeeeeee-Nnnnnnnnnn-Uuuuuuuuuuu"},
		{"NyffsGeyylB", "N-Yy-Fff-Ffff-Sssss-Gggggg-Eeeeeee-Yyyyyyyy-Yyyyyyyyy-Llllllllll-Bbbbbbbbbbb"},
		{"MjtkuBovqrU", "M-Jj-Ttt-Kkkk-Uuuuu-Bbbbbb-Ooooooo-Vvvvvvvv-Qqqqqqqqq-Rrrrrrrrrr-Uuuuuuuuuuu"},
		{"EvidjUnokmM", "E-Vv-Iii-Dddd-Jjjjj-Uuuuuu-Nnnnnnn-Oooooooo-Kkkkkkkkk-Mmmmmmmmmm-Mmmmmmmmmmm"},
		{"HbideVbxncC", "H-Bb-Iii-Dddd-Eeeee-Vvvvvv-Bbbbbbb-Xxxxxxxx-Nnnnnnnnn-Cccccccccc-Ccccccccccc"},
	}

	for _, test := range testCases {
		got := Accum(test.s)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = Accum2(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = Accum3(test.s)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
