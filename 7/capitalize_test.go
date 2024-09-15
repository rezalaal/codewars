package kata7

import "testing"

func TestCapitalize(t *testing.T) {
	testCases := []struct {
		st  string
		out []string
	}{
		{"abcdef", []string{"AbCdEf", "aBcDeF"}},
		{"codewars", []string{"CoDeWaRs", "cOdEwArS"}},
		{"abracadabra", []string{"AbRaCaDaBrA", "aBrAcAdAbRa"}},
		{"codewarriors", []string{"CoDeWaRrIoRs", "cOdEwArRiOrS"}},
		{"indexinglessons", []string{"InDeXiNgLeSsOnS", "iNdExInGlEsSoNs"}},
		{"codingisafunactivity", []string{"CoDiNgIsAfUnAcTiViTy", "cOdInGiSaFuNaCtIvItY"}},
	}

	for _, test := range testCases {
		got := Capitalize(test.st)
		want := test.out

		assertCorrectMessage(t, want, got)
	}
}
