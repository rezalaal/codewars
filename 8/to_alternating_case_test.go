package kata8

import "testing"

func TestToAlternatingCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"HELLO WORLD", "hello world"},
		{"hello WORLD", "HELLO world"},
		{"HeLLo WoRLD", "hEllO wOrld"},
		{"12345", "12345"},
		{"1a2b3c4d5e", "1A2B3C4D5E"},
		{"String.prototype.toAlternatingCase", "sTRING.PROTOTYPE.TOaLTERNATINGcASE"},
		{"Hello World", "hELLO wORLD"},
		{"altERnaTIng cAsE", "ALTerNAtiNG CaSe"},
		{"ALTerNAtiNG CaSe", "altERnaTIng cAsE"},
		{"altERnaTIng cAsE <=> ALTerNAtiNG CaSe", "ALTerNAtiNG CaSe <=> altERnaTIng cAsE"},
		{"ALTerNAtiNG CaSe <=> altERnaTIng cAsE", "altERnaTIng cAsE <=> ALTerNAtiNG CaSe"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			got1 := ToAlternatingCase(test.input)
			if got1 != test.expected {
				t.Errorf("ToAlternatingCase: expected %v but got %v", test.expected, got1)
			}

			got2 := ToAlternatingCase2(test.input)
			if got2 != test.expected {
				t.Errorf("ToAlternatingCase2: expected %v but got %v", test.expected, got2)
			}
		})
	}
}
