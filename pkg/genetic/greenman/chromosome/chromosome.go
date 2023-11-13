package chromosome

import "github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/dnastrand"

type Chromosome struct {
	Father dnastrand.DNAStrand
	Mother dnastrand.DNAStrand
}
