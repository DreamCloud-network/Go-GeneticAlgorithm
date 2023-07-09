package dna

type Codon byte

// Codons definition
const (
	Dominant Codon = iota
	Recessive
	Start
	End
)
