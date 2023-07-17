package dna

import (
	"strconv"
)

type Codon byte

type Gene struct {
	Code []Codon
}

func NewGene() Gene {
	return Gene{
		Code: make([]Codon, 0),
	}
}

func (gene *Gene) String() string {
	str := ""

	for _, codon := range gene.Code {
		str += strconv.Itoa(int(codon))
	}

	return str
}

func (gene *Gene) Duplicate() *Gene {
	newGene := Gene{
		Code: make([]Codon, len(gene.Code)),
	}

	copy(newGene.Code, gene.Code)

	return &newGene
}
