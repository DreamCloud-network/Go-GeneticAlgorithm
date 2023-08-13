package codons

import (
	"log"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

// Helps to decoder the codons.
type Decoder struct {
	//INIT_CODONSearchState   int
	//END_CODONSearchState    int
	//CHIASM_CODONSearchState int

	actualCodon Codon
}

func NewDecoder() Decoder {
	return Decoder{
		actualCodon: EMPTY_CODON,
	}
}

func (d *Decoder) NewFeda(feda fedas.Feda) error {
	if !feda.IsValid() {
		log.Println("codons.Decoder.NewFeda - Invalid feda.")
		return ErrorInvalidFeda
	}

	d.actualCodon[0] = d.actualCodon[1]
	d.actualCodon[1] = d.actualCodon[2]
	d.actualCodon[2] = feda

	return nil
}

// Returns true when a END_CODON is found.
func (d *Decoder) INIT_CODON() bool {
	return d.actualCodon == INIT_CODON
}

// Returns true when a END_CODON is found.
func (d *Decoder) END_CODON() bool {
	return d.actualCodon == END_CODON
}

// Returns true when a EMPTY_CODON is found.
func (d *Decoder) EMPTY_CODON() bool {
	return d.actualCodon == EMPTY_CODON
}

// Returns true when a CHIASM_CODON is found.
func (d *Decoder) CHIASM_CODON() bool {
	return d.actualCodon == CHIASM_CODON
}

// Reset the decoder state.
func (d *Decoder) Reset() {
	d.actualCodon = EMPTY_CODON
}
