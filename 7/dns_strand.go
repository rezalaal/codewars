package kata7


import (
	"strings"
)

/*
#Complementary DNA

Deoxyribonucleic acid (DNA) is a chemical found in the nucleus of cells and carries the "instructions" for the development and functioning of living organisms.

If you want to know more: http://en.wikipedia.org/wiki/DNA

In DNA strings, symbols "A" and "T" are complements of each other, as "C" and "G". Your function receives one side of the DNA (string, except for Haskell); you need to return the other complementary side. DNA strand is never empty or there is no DNA at all (again, except for Haskell).

More similar exercise are found here: http://rosalind.info/problems/list-view/ (source)

Example: (input --> output)

"ATTGC" --> "TAACG"
"GTAT" --> "CATA"
*/
func DNAStrand(dna string) string {
	res := ""
	for _, char := range dna {
		switch string(char) {
		case "A":
			res += "T"
		case "T":
			res += "A"
		case "G":
			res += "C"
		case "C":
			res += "G"
		default:
			res += string(char)
		}
	}
	return res
}

var dnaReplacer *strings.Replacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"C", "G",
	"G", "C",
  )
  
func DNAStrand1(dna string) string {
	return dnaReplacer.Replace(dna)
}

func DNAStrand2(dna string) string {
	replacer := strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G")
	  return(replacer.Replace(dna))
}

func DNAStrand3(dna string) string {
	return strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G").Replace(dna)
}