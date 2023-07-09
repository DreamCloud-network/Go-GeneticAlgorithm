package dna

import (
	"strconv"
)

type Gene struct {
	Code []Codon
}

func NewGene(dominant bool) *Gene {
	newGene := Gene{
		Code: make([]Codon, 0),
	}

	return &newGene
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
