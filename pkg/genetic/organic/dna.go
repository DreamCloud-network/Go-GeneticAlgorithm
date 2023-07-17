package organic

type NucleotideBase byte

const (
	A NucleotideBase = iota
	T
	C
	G
	U
)

type Nucleotide struct {
	Base NucleotideBase
}

type Codon struct {
	Bases [3]Nucleotide
}

type Gene struct {
	Codons []Codon
}

type Strand struct {
	Genes []Gene
}

type SingleHelixDNA struct {
	Strands Strand
}

type DoubleHelixDNA struct {
	Strands [2]Strand
}
