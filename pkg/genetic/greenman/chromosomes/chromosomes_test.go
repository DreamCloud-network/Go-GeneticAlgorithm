package chromosomes

import (
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

func TestCrossoverReposition(t *testing.T) {
	t.Log("TestCrossoverReposition")

	chromosomeLenght := 100

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	c1ChiasmPositions := make([]int, r1.Intn(chromosomeLenght/2)+1)

	for i := range c1ChiasmPositions {
		c1ChiasmPositions[i] = r1.Intn(chromosomeLenght/len(c1ChiasmPositions)) + 1
		if i > 0 {
			c1ChiasmPositions[i] += c1ChiasmPositions[i-1]
		}
	}

	c2ChiasmPositions := make([]int, r1.Intn(chromosomeLenght/2)+1)

	for i := range c2ChiasmPositions {
		c2ChiasmPositions[i] = r1.Intn(chromosomeLenght/len(c2ChiasmPositions)) + 1
		if i > 0 {
			c2ChiasmPositions[i] += c2ChiasmPositions[i-1]
		}
	}

	log.Println("c1ChiasmPositions lenght: ", len(c1ChiasmPositions))
	log.Println("c1ChiasmPositions: ", c1ChiasmPositions)
	log.Println("c2ChiasmPositions lenght: ", len(c2ChiasmPositions))
	log.Println("c2ChiasmPositions: ", c2ChiasmPositions)

	finalPosition := crossoverRegulation(c1ChiasmPositions, c2ChiasmPositions, chromosomeLenght)
	log.Println("finalPosition lenght: ", len(finalPosition))
	log.Println("finalPosition: ", finalPosition)
}

func TestCrossover(t *testing.T) {
	t.Log("TestCrossover")

	motherChromosome := NewChromosome()
	fatherChromosome := NewChromosome()

	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	genesCount := r1.Intn(5) + 5

	for i := 0; i < genesCount; i++ {
		motherGeneBuilder := genes.NewBuilder()

		if (r1.Intn(100) < 50) && (i > 0) {
			motherChromosome.AddCodon(codons.CHIASM_CODON)
		}

		for j := 0; j < (r1.Intn(5) + 1); j++ {
			newCodon := codons.NewBuilder()
			newCodon.GenerateRandomCodon()
			motherGeneBuilder.AddCodon(*newCodon.GetCodon())
		}
		motherGene := motherGeneBuilder.GetGene()
		motherChromosome.AddGene(motherGene)

		if (r1.Intn(100) < 50) && (i > 0) {
			fatherChromosome.AddCodon(codons.CHIASM_CODON)
		}

		fatherGeneBuilder := genes.NewBuilder()
		for j := 0; j < (r1.Intn(5) + 1); j++ {
			newCodon := codons.NewBuilder()
			newCodon.GenerateRandomCodon()
			fatherGeneBuilder.AddCodon(*newCodon.GetCodon())
		}
		fatherGene := fatherGeneBuilder.GetGene()
		fatherChromosome.AddGene(fatherGene)
	}

	log.Println("motherChromosome: ", motherChromosome.String())
	log.Println("fatherChromosome: ", fatherChromosome.String())

	motherChromosome.Crossover(&fatherChromosome)

	log.Println("Croosover motherChromosome: ", motherChromosome.String())
	log.Println("Crossover fatherChromosome: ", fatherChromosome.String())

}
