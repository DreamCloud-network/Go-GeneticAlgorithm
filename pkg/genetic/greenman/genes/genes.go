package genes

import (
	"strings"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

type Gene struct {
	code []codons.Codon
}

// Returna a new blank gene.
func NewGene() Gene {
	newGene := Gene{
		code: make([]codons.Codon, 0),
	}
	return newGene
}

func (gene *Gene) String() string {
	var geneStr strings.Builder

	for _, codon := range gene.code {
		geneStr.WriteString(codon.String())
	}

	return geneStr.String()
}

// Return all the fedas of the gene.
func (gene *Gene) GetFedas() []fedas.Feda {

	fedas := make([]fedas.Feda, 0, len(gene.code)*3)

	for _, codon := range gene.code {
		fedas = append(fedas, codon.GetFedas()...)
	}

	return fedas
}

/*
// Reset code to a new initiaded code. >
func (gene *Gene) ResetCode() {
	gene.code = NewGene().code
}

// Return the code of the gene
func (gene *Gene) GetCodes() []fedas.Feda {
	return gene.code
}

// Return one specific code of the gene
func (gene *Gene) ReadCode(codePosition int) (fedas.Feda, error) {
	if codePosition > len(gene.code) {
		return fedas.Peith, ErrOutOfRange
	}
	return gene.code[codePosition], nil
}

func (gene *Gene) AppendFeda(feda fedas.Feda) error {
	if feda == fedas.INITIATOR {
		return ErrNotPermited
	}

	gene.code = append(gene.code, feda)

	return nil
}

func (gene *Gene) RemoveLastFeda() fedas.Feda {
	removedFeda, _ := gene.RemoveFeda(len(gene.code) - 1)
	return removedFeda
}

func (gene *Gene) RemoveFirstFeda() fedas.Feda {
	removedFeda, _ := gene.RemoveFeda(1)
	return removedFeda
}

// Remove any feda from the gene
func (gene *Gene) RemoveFeda(index int) (fedas.Feda, error) {
	if index > len(gene.code) {
		return fedas.Peith, ErrOutOfRange
	}

	if index == 0 {
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

func (gene *Gene) SetCode(codePosition int, feda fedas.Feda) error {
	if codePosition > len(gene.code) {
		return ErrOutOfRange
	}

	if codePosition == 0 {
		return ErrNotPermited
	}

	gene.code[codePosition] = feda
	return nil
}
*/
