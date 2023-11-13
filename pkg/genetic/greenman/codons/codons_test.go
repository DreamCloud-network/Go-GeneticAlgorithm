package codons

import (
	"testing"

	"github.com/google/uuid"
)

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

func TestUUIDCodons(t *testing.T) {
	t.Log("TestUUIDCodons")

	newUUID := uuid.New().String()

	codons, err := UUIDToCodons(newUUID)

	if err != nil {
		t.Error(err)
		return
	}

	codonStr := ""
	for _, codon := range codons {
		codonStr += codon.String()
	}

	t.Log("UUID: ", newUUID, "\n\r-> ", codonStr)

	t.Log("Converting Codons to UUID")

	newUUID2, err := CodonsToUUID(codons)

	if err != nil {
		t.Error(err)
		return
	}

	if newUUID != newUUID2 {
		t.Errorf("Error - UUID: %s -> %s", newUUID, newUUID2)
		return
	}

	t.Log("Original: ", newUUID, "\n\rAfter Convertion: ", newUUID2)

}
