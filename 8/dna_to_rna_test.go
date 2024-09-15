package kata8

import "testing"

func TestDNAtoRNA(t *testing.T) {
	got := DNAtoRNA("GCAT")
	expected := "GCAU"

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}

	got = DNAtoRNA1("GCAT")
	expected = "GCAU"

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}

	got = DNAtoRNA2("GCAT")
	expected = "GCAU"

	if got != expected {
		t.Errorf("Expected %v but got %v", expected, got)
	}
}
