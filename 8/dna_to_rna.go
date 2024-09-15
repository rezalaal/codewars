package kata8

import "strings"

/*#DNA to RNA Conversion

Deoxyribonucleic acid, DNA is the primary information storage molecule in biological systems. It is composed of four nucleic acid bases Guanine ('G'), Cytosine ('C'), Adenine ('A'), and Thymine ('T').

Ribonucleic acid, RNA, is the primary messenger molecule in cells. RNA differs slightly from DNA its chemical structure and contains no Thymine. In RNA Thymine is replaced by another nucleic acid Uracil ('U').

Create a function which translates a given DNA string into RNA.

For example:

"GCAT"  =>  "GCAU"
The input string can be of arbitrary length - in particular, it may be empty. All input is guaranteed to be valid, i.e. each input string will only ever consist of 'G', 'C', 'A' and/or 'T'.

*/
func DNAtoRNA1(dna string) string {
	var rna strings.Builder
	for _,v := range dna {
		if (v == 'T') {
			rna.WriteRune('U')
		}else{
			rna.WriteRune(v)
		}
	}

	return rna.String()
}

func DNAtoRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func DNAtoRNA2(dna string) string {
	var rna string

	for _, letter := range dna{
		if letter == 'T'{
			letter = 'U'
		}
		rna+=string(letter)
	}

	return rna
}