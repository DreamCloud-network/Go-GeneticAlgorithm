package codons

import (
	"strings"

	"github.com/GreenMan-Network/Go-GeneticAlgorithm/pkg/genetic/greenman/fedas"
)

type Codon [3]fedas.Feda

var INIT_CODON Codon = [3]fedas.Feda{fedas.INITIATOR, fedas.SPACE, fedas.SPACE}
var END_CODON Codon = [3]fedas.Feda{fedas.SPACE, fedas.SPACE, fedas.TERMINATOR}
var EMPTY_CODON Codon = [3]fedas.Feda{fedas.SPACE, fedas.SPACE, fedas.SPACE}

var CHIASM_CODON Codon = [3]fedas.Feda{fedas.Ebad, fedas.Oir, fedas.Uillenn}

// Returns the string representation of the codon.
func (c *Codon) String() string {
	var codonStr strings.Builder

	for _, feda := range c {
		codonStr.WriteString(feda.String())
	}

	return codonStr.String()
}

// Return the fedas of the codon.
func (c *Codon) GetFedas() []fedas.Feda {
	return c[:]
}
