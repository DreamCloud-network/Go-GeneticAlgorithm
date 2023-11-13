package genes

import (
	"strings"
	"sync"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/codons"
	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
)

type Gene struct {
	sync.Mutex

	preCode []codons.Codon
	code    []codons.Codon
}

// Returna a new blank gene.
func NewGene() *Gene {
	newGene := &Gene{
		Mutex: sync.Mutex{},

		preCode: make([]codons.Codon, 0),
		code:    make([]codons.Codon, 1),
	}

	newGene.code[0] = codons.INIT_CODON

	return newGene
}

func (gene *Gene) String() string {
	gene.Lock()
	defer gene.Unlock()

	var geneStr strings.Builder

	for _, codon := range gene.preCode {
		geneStr.WriteString(codon.String())
	}

	geneStr.WriteString(" ")

	for _, codon := range gene.code {
		geneStr.WriteString(codon.String())
	}

	return geneStr.String()
}

// Return all the feda of the gene.
func (gene *Gene) ReadFeda() []feda.Fid {
	gene.Lock()
	defer gene.Unlock()

	feda := make([]feda.Fid, ((len(gene.code) + len(gene.preCode)) * 3))

	for _, codon := range gene.preCode {
		feda = append(feda, codon.ToFeda()...)
	}

	for _, codon := range gene.code {
		feda = append(feda, codon.ToFeda()...)
	}

	return feda
}

// Return the codons of the gene pre code region.
func (gene *Gene) ReadPreCode() []codons.Codon {
	gene.Lock()
	defer gene.Unlock()

	precodeCodons := make([]codons.Codon, len(gene.preCode))
	copy(precodeCodons, gene.preCode)
	return precodeCodons
}

// Return the codons of the gene code region.
func (gene *Gene) ReadCode() []codons.Codon {
	gene.Lock()
	defer gene.Unlock()

	codeCodons := make([]codons.Codon, len(gene.code))
	copy(codeCodons, gene.code)
	return codeCodons
}

// Return the codons of the gene.
func (gene *Gene) ReadCodons() []codons.Codon {
	//gene.Lock()
	//defer gene.Unlock()

	codonsCode := make([]codons.Codon, 0, len(gene.code)+len(gene.preCode))

	precode := gene.ReadPreCode()
	code := gene.ReadCode()

	codonsCode = append(codonsCode, precode...)
	codonsCode = append(codonsCode, code...)

	return codonsCode
}

// Returns a exact copy of the chromosome.
func (gene *Gene) Duplicate() *Gene {
	gene.Lock()
	defer gene.Unlock()

	newGene := NewGene()

	newGene.preCode = make([]codons.Codon, len(gene.preCode))
	copy(newGene.preCode, gene.preCode)

	newGene.code = make([]codons.Codon, len(gene.code))
	copy(newGene.code, gene.code)

	return newGene
}

// Return the codons of the gene code region excluding initator.
func (gene *Gene) GetRawCode() []codons.Codon {
	gene.Lock()
	defer gene.Unlock()

	rawCodeCodons := make([]codons.Codon, len(gene.code)-1)
	copy(rawCodeCodons, gene.code[1:len(gene.code)])
	return rawCodeCodons
}

// Return true if the gene contains a chiasm codon.
func (gene *Gene) HasChiasm() bool {
	gene.Lock()
	defer gene.Unlock()

	if (gene.preCode == nil) || (len(gene.preCode) <= 0) {
		return false
	}

	return gene.preCode[0].IsChiasm()
}

// Remove all the chiasm codons from the gene.
func (gene *Gene) RemoveChiasm() {

	if gene.HasChiasm() {
		gene.Lock()
		defer gene.Unlock()

		if len(gene.preCode) > 1 {
			gene.preCode = gene.preCode[1:]
		} else {
			gene.preCode = make([]codons.Codon, 0)
		}
	}
}

// Add a quiasm code in the pre code region of the gene.
// Return an error if the gene is finalized.
// Do nothing if there already is a chiasm code.
func (gene *Gene) AddChiasm() {
	if !gene.HasChiasm() {
		gene.Lock()
		defer gene.Unlock()

		// Add the chiasm always in the first position.
		gene.preCode = append([]codons.Codon{codons.CHIASM_CODON}, gene.preCode...)
	}
}

// Return true if the gene is disabled by a pre DISABLED_CODON.
func (gene *Gene) IsEnabled() bool {
	gene.Lock()
	defer gene.Unlock()

	for _, codon := range gene.preCode {
		if codon.IsEnable() {
			return true
		}
	}

	return false
}

// Disable the gene removing the ENABLE_CODON if it is present.
func (gene *Gene) Disable() {
	// confirm if the ENABLE_CODON is not already in the gene
	if gene.IsEnabled() {
		gene.Lock()
		defer gene.Unlock()

		for i, codon := range gene.preCode {
			if codon.IsEnable() {
				gene.preCode = append(gene.preCode[:i], gene.preCode[i+1:]...)
			}
		}
	}
}

// Enable a codon adding the ENABLE_CODON is it is not already in the pre code.
func (gene *Gene) Enable() {
	if !gene.IsEnabled() {
		gene.Lock()
		defer gene.Unlock()

		gene.preCode = append(gene.preCode, codons.ENABLE_CODON)
	}
}

// Add a new valid codon to the gene.
// Doens't allow to add INIT_CODON or END_CODON.
func (gene *Gene) AddCodon(codon codons.Codon) error {
	if !codon.AreFeda() && !codon.IsEmpty() {
		return ErrNotPermited
	}

	gene.Lock()
	defer gene.Unlock()

	gene.code = append(gene.code, codon)

	return nil
}

// Add valid codons to the gene.
// Doens't allow to add INIT_CODON or END_CODON.
func (gene *Gene) AddCodons(codons []codons.Codon) error {
	for _, codon := range codons {
		if !codon.AreFeda() && !codon.IsEmpty() {
			return ErrNotPermited
		}
	}

	gene.Lock()
	defer gene.Unlock()

	gene.code = append(gene.code, codons...)

	return nil
}
