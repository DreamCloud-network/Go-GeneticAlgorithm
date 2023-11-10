package codons

import "testing"

func TestUintToCodons(t *testing.T) {
	t.Log("TestUintToCodons")

	for val := uint(159000); val <= 160001; val++ {
		codons := UintToCodons(val)

		codonStr := ""
		for _, codon := range codons {
			codonStr += codon.String()
		}

		newVal, err := CodonsToUint(codons)
		if err != nil {
			t.Error(err)
			return
		}

		if val != newVal {
			t.Errorf("Error - Value: %d -> %s -> %d", val, codonStr, newVal)
			return
		}

		t.Log("Value: ", val, " -> ", codonStr, " -> ", newVal)
	}
}
