package kata7

import "testing"

func TestSimpleStringReversalII(t *testing.T) {

	testCases := []struct {
		s string
		a int
		b int
		out string
	}{
		{"codewars",1,5,"cawedors"},
		{"codingIsFun",2,100, "conuFsIgnid"},
		{"FunctionalProgramming", 2, 15,"FuargorPlanoitcnmming"},
		{"abcefghijklmnopqrstuvwxyz",0,20, "vutsrqponmlkjihgfecbawxyz"},
		{"abcefghijklmnopqrstuvwxyz",5,20, "abcefvutsrqponmlkjihgwxyz"},
	}

	for _, test := range testCases {
		got := SimpleStringReversalII(test.s, test.a, test.b)
		want := test.out
		assertCorrectMessage(t, want, got)
	}
	  
}