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

type Allele struct {
	Genes []Gene
}

type Strand struct {
	Alleles []Allele
}

type DNA struct {
	Strands [2]Strand
}

type Chromosomes struct {
	Father *DNA
	Mother *DNA
}
