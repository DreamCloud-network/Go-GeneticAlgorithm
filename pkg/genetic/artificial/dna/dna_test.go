package dna

import (
	"log"
	"testing"
)

func TestDNADuplication(t *testing.T) {
	t.Log("TestDNADuplication")

	dna := NewDNA()
	dna.Genes = append(dna.Genes, NewGene())
	dna.Genes[0].Code = append(dna.Genes[0].Code, Codon(1))

	log.Println("DNA1: ", dna)

	dna2 := *dna.Duplicate()
	log.Println("DNA2: ", dna2)

	dna2.Genes[0].Code[0] = Codon(2)
	log.Println("DNA1: ", dna)
	log.Println("DNA2: ", dna2)

	dna.Genes[0].Code[0] = Codon(3)
	log.Println("DNA1: ", dna)
	log.Println("DNA2: ", dna2)
}
