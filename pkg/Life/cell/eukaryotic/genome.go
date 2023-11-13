package eukaryotic

import (
	"log"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/genes"
)

type genesRegulations int

const (
	ENERGY_CONVERSION_RATE genesRegulations = iota
)

type Genome struct {
	regulations dnastrand.DNAStrand
}

func newGenome() Genome {
	newGenome := Genome{
		regulations: dnastrand.NewDNAStrand(),
	}

	newGenome.newRegulationGene()

	return newGenome
}

// Generate a new basic regulation gene.
func (newGene *Genome) newRegulationGene() {
	// ENERGY_CONVERSION_RATE
	converionRateCodons := codons.UintToCodons(uint(4))
	genesBuilder := genes.NewBuilder()
	for _, codon := range converionRateCodons {
		genesBuilder.AddCode(codon)
	}
	converionRateGene := genesBuilder.Finish()
	newGene.regulations.AddGene(*converionRateGene)
}

// Return the conversion rate of the mitochondria.
func (genome Genome) getConversionRate() (int, error) {
	conversionRateCodons := genome.regulations[ENERGY_CONVERSION_RATE].GetRawCode()
	converionRate, err := codons.CodonsToUint(conversionRateCodons)
	if err != nil {
		log.Println("mitochondria.Genome.getConversionRate - error converting codons to uint")
		return 0, err
	}

	return int(converionRate), nil
}
