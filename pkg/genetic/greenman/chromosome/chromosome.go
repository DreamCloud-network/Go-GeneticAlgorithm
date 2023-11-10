package chromosome

import "github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"

type Chromosome struct {
	Father dnastrand.DNAStrand
	Mother dnastrand.DNAStrand
}
