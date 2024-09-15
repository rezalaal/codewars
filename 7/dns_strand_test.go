package kata7

import "testing"

func TestDNAStrand(t *testing.T) {
	testCases := []struct {
		dna string
		out string
	}{
		{"AAAA", "TTTT"},
		{"ATTGC", "TAACG"},
		{"GTAT", "CATA"},
	}

	for _, test := range testCases {
		got := DNAStrand(test.dna)
		want := test.out
		assertCorrectMessage(t, want, got)

		got = DNAStrand1(test.dna)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = DNAStrand2(test.dna)
		want = test.out
		assertCorrectMessage(t, want, got)

		got = DNAStrand3(test.dna)
		want = test.out
		assertCorrectMessage(t, want, got)
	}
}
