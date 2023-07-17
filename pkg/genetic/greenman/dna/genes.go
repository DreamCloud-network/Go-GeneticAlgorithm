package dna

import (
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/dna/fedas"
)

type Gene struct {
	code []fedas.Feda
}

func NewGene() Gene {
	newGene := Gene{
		code: make([]fedas.Feda, 1),
	}
	newGene.code[0] = fedas.INI

	return newGene
}

// Reset code to a new initiaded code. >
func (gene *Gene) ResetCode() {
	newGene := NewGene()
	gene.code = newGene.code
}

// Return the code of the gene
func (gene *Gene) GetCode() []fedas.Feda {
	return gene.code
}

func (gene *Gene) String() string {
	str := ""

	for _, feda := range gene.code {
		str += feda.String()
	}

	return str
}

func (gene *Gene) AppendFeda(feda fedas.Feda) {
	if feda != fedas.INI {
		gene.code = append(gene.code, feda)
	}
}

func (gene *Gene) RemoveLastFeda() fedas.Feda {
	removedFeda, _ := gene.RemoveFeda(len(gene.code) - 1)
	return removedFeda
}

func (gene *Gene) RemoveFirstFeda() fedas.Feda {
	removedFeda, _ := gene.RemoveFeda(1)
	return removedFeda
}

// Remove any feda from the gene, except the > (INI)
func (gene *Gene) RemoveFeda(index int) (fedas.Feda, error) {
	if index > len(gene.code) {
		return fedas.Peith, ErrOutOfRange
	}

	if index <= 0 {
		return fedas.Peith, ErrNotPermited
	}

	removedFeda := gene.code[index]

	gene.code = append(gene.code[:index], gene.code[index+1:]...)
	return removedFeda, nil
}

func (gene *Gene) Duplicate() *Gene {
	newGene := Gene{
		code: make([]fedas.Feda, len(gene.code)),
	}

	copy(newGene.code, gene.code)

	return &newGene
}
