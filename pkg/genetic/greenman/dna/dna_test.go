package dna

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/dna/fedas"
)

func TestDNADuplication(t *testing.T) {
	t.Log("TestDNADuplication")

	dna := NewDNA()
	dna.Genes = append(dna.Genes, NewGene())
	dna.Genes[0].code = append(dna.Genes[0].code, fedas.Feda(0))

	log.Println("DNA1: ", dna)

	dna2 := *dna.Duplicate()
	log.Println("DNA2: ", dna2)

	dna2.Genes[0].code[1] = fedas.Feda(1)
	log.Println("DNA1: ", dna)
	log.Println("DNA2: ", dna2)

	dna.Genes[0].code[1] = fedas.Feda(2)
	log.Println("DNA1: ", dna)
	log.Println("DNA2: ", dna2)
}

func TestString(t *testing.T) {
	t.Log("TestString")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	numGenes := r1.Intn(5) + 1

	dna := NewDNA()

	for i := 0; i < numGenes; i++ {
		dna.Genes = append(dna.Genes, NewGene())

		for j := 0; j < (r1.Intn(10) + 1); j++ {
			dna.Genes[i].AppendFeda(fedas.Feda(r1.Intn(int(fedas.Peith))))
		}
	}

	log.Println("DNA: ", dna.String())

	test := 'ᛒ'
	log.Println("test: ", string(test))
}
