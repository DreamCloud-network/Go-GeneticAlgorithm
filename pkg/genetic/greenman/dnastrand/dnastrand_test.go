package dnastrand

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

/*
	func TestCrossoverReposition(t *testing.T) {
		t.Log("TestCrossoverReposition")

		dnaSTrandLenght := 100

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

		c1ChiasmPositions := make([]int, r1.Intn(dnaSTrandLenght/2)+1)

		for i := range c1ChiasmPositions {
			c1ChiasmPositions[i] = r1.Intn(dnaSTrandLenght/len(c1ChiasmPositions)) + 1
			if i > 0 {
				c1ChiasmPositions[i] += c1ChiasmPositions[i-1]
			}
		}

		c2ChiasmPositions := make([]int, r1.Intn(dnaSTrandLenght/2)+1)

		for i := range c2ChiasmPositions {
			c2ChiasmPositions[i] = r1.Intn(dnaSTrandLenght/len(c2ChiasmPositions)) + 1
			if i > 0 {
				c2ChiasmPositions[i] += c2ChiasmPositions[i-1]
			}
		}

		log.Println("c1ChiasmPositions lenght: ", len(c1ChiasmPositions))
		log.Println("c1ChiasmPositions: ", c1ChiasmPositions)
		log.Println("c2ChiasmPositions lenght: ", len(c2ChiasmPositions))
		log.Println("c2ChiasmPositions: ", c2ChiasmPositions)

		finalPosition := crossoverRegulation(c1ChiasmPositions, c2ChiasmPositions, dnaSTrandLenght)
		log.Println("finalPosition lenght: ", len(finalPosition))
		log.Println("finalPosition: ", finalPosition)
	}

	func TestCrossover(t *testing.T) {
		t.Log("TestCrossover")

		motherDNAStrand := NewDNAStrand()
		fatherDNAStrand := NewDNAStrand()

		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

		genesCount := r1.Intn(5) + 5

		for i := 0; i < genesCount; i++ {
			motherGene := genes.NewGene()
			for j := 0; j < (r1.Intn(5) + 1); j++ {
				newCodonBuilder := codons.NewBuilder()
				newCodonBuilder.GenerateRandomFedaCodon()
				motherGene.AddCodon(newCodonBuilder.GetCodon())
			}

			//if r1.Intn(100) < 50 {
			motherGene.AddChiasm()
			//}
			motherDNAStrand.AddGene(motherGene)

			fatherGene := genes.NewGene()
			for j := 0; j < (r1.Intn(5) + 1); j++ {
				newCodonBuilder := codons.NewBuilder()
				newCodonBuilder.GenerateRandomFedaCodon()
				fatherGene.AddCodon(newCodonBuilder.GetCodon())
			}

			//if r1.Intn(100) < 50 {
			fatherGene.AddChiasm()
			//}
			fatherDNAStrand.AddGene(fatherGene)
		}

		log.Println("motherChromosome: ", motherDNAStrand.String())
		log.Println("fatherChromosome: ", fatherDNAStrand.String())

		motherDNAStrand.Crossover(&fatherDNAStrand)

		log.Println("Croosover motherChromosome: ", motherDNAStrand.String())
		log.Println("Crossover fatherChromosome: ", fatherDNAStrand.String())
	}
*/
func TestDuplicate(t *testing.T) {
	t.Log("TestDuplicate")

	dnaStrand := NewDNAStrand()

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	genesCount := r1.Intn(5) + 5

	for i := 0; i < genesCount; i++ {
		newGene := genes.NewGene()
		for j := 0; j < (r1.Intn(5) + 1); j++ {
			newCodonBuilder := codons.NewBuilder()
			newCodonBuilder.GenerateRandomFedaCodon()
			newGene.AddCodon(newCodonBuilder.GetCodon())
		}

		if r1.Intn(100) < 50 {
			newGene.AddChiasm()
		}
		dnaStrand.AddGene(newGene)
	}

	dnaStrandCopy := dnaStrand.Duplicate()

	log.Println("original:\n", dnaStrand.String())
	log.Println("copy:\n", dnaStrandCopy.String())

	newGene := genes.NewGene()
	for j := 0; j < (r1.Intn(5) + 1); j++ {
		newCodonBuilder := codons.NewBuilder()
		newCodonBuilder.GenerateRandomFedaCodon()
		newGene.AddCodon(newCodonBuilder.GetCodon())
	}

	if r1.Intn(100) < 50 {
		newGene.AddChiasm()
	}
	dnaStrandCopy.AddGene(newGene)

	log.Println("New original:\n", dnaStrand.String())
	log.Println("New copy:\n", dnaStrandCopy.String())
}
