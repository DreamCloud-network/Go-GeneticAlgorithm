package genes

import "github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"

// Helps to build a new gene by adding one codon or one feda at a time.
// Also helps to check if the gene is complete and use only the values that corresponds to the codons and fedas.
type Builder struct {
	gene      Gene
	finalized bool
}

// Returns a new builder already initiaded.
func NewBuilder() Builder {
	newBuilder := Builder{
		gene:      NewGene(),
		finalized: false,
	}

	newBuilder.gene.code = append(newBuilder.gene.code, codons.INIT_CODON)

	return newBuilder
}

// Add a new valid codon to the gene.
// Doens't allow to add INIT_CODON or END_CODON.
func (b *Builder) AddCodon(codon codons.Codon) error {
	if codon == codons.INIT_CODON || codon == codons.END_CODON || b.finalized {
		return ErrNotPermited
	}

	b.gene.code = append(b.gene.code, codon)

	return nil
}

// Returns true if the gene is complete, otherwise returns false.
func (b *Builder) IsFinalized() bool {
	return b.finalized
}

// Finalize and return the gene.
// After finalized the gene can't be modified.
func (b *Builder) GetGene() *Gene {
	// In this case, the gene has only the initiator codon.
	if len(b.gene.code) <= 1 {
		return nil
	}

	if !b.IsFinalized() {
		b.gene.code = append(b.gene.code, codons.END_CODON)
		b.finalized = true
	}

	return &b.gene
}
