package codons

import (
	"log"
	"strings"

	"github.com/DreamCloud-network/Go-GeneticAlgorithm/pkg/genetic/greenman/feda"
)

type Codon [3]feda.Fid

// Inside gene codons
var INIT_CODON Codon = [3]feda.Fid{feda.INITIATOR, feda.SPACE, feda.SPACE}

// var END_CODON Codon = [3]feda.Fid{feda.SPACE, feda.SPACE, feda.TERMINATOR}

var EMPTY_CODON Codon = [3]feda.Fid{feda.SPACE, feda.SPACE, feda.SPACE}

// Out of gene codons
var CHIASM_CODON Codon = [3]feda.Fid{feda.Ebad, feda.Oir, feda.Uillenn}
var ENABLE_CODON Codon = [3]feda.Fid{feda.Oir, feda.Oir, feda.Oir}

// Returns the string representation of the codon.
func (c Codon) String() string {
	var codonStr strings.Builder

	for _, feda := range c {
		codonStr.WriteString(feda.String())
	}

	return codonStr.String()
}

// Return the fedas of the codon.
func (c Codon) ToFeda() []feda.Fid {
	return c[:]
}

// Return the fedas of the codon.
func (c Codon) ReadFeda() []feda.Fid {
	feda := make([]feda.Fid, 3)
	copy(feda, c[:])
	return feda
}

// Return true if the codon is composed only by forfeda.
func (c Codon) AreForfeda() bool {
	for _, fid := range c {
		if !fid.IsForfid() {
			return false
		}
	}
	return true
}

// Return true if the codon is composed only by feda.
func (c Codon) AreFeda() bool {
	for _, fid := range c {
		if !fid.IsFid() && (fid != feda.SPACE) {
			return false
		}
	}
	return true
}

// Return true if it is an initiator codon.
func (c Codon) IsInitiator() bool {
	return c == INIT_CODON
}

// Return true if it is an end codon.
/*
func (c Codon) IsEnd() bool {
	return c == END_CODON
}
*/
// Return true if it is an empty codon.
func (c Codon) IsEmpty() bool {
	return c == EMPTY_CODON
}

// Return true if it is a chiasm codon.
func (c Codon) IsChiasm() bool {
	return c == CHIASM_CODON
}

// Return true if it is a enable codon.
func (c Codon) IsEnable() bool {
	return c == ENABLE_CODON
}

// Return the codons representation of the uint.
// The most significant value is to the left.
// Beith (ᚁ) = 0.
func UintToCodons(val uint) []Codon {
	var codons []Codon

	valFeda := feda.UintToFeda(val)

	newCodonBuilder := NewBuilder()
	newCodonBuilder.SetReverseBuilder(true)

	for pos := len(valFeda) - 1; pos >= 0; pos-- {
		newCodonBuilder.AddFid(valFeda[pos])
		if newCodonBuilder.IsComplete() {
			codons = append([]Codon{newCodonBuilder.GetCodon()}, codons...)
			newCodonBuilder = NewBuilder()
			newCodonBuilder.SetReverseBuilder(true)
		}
	}

	if newCodonBuilder.HasFid() {
		for !newCodonBuilder.IsComplete() {
			newCodonBuilder.AddFid(feda.SPACE)
		}

		codons = append([]Codon{newCodonBuilder.GetCodon()}, codons...)
	}

	return codons
}

// Return the uint representation of the codons.
// The most significant value is to the left.
// Beith (ᚁ) = 0.
func CodonsToUint(codons []Codon) (uint, error) {
	valFeda := make([]feda.Fid, 0, len(codons)*3)

	for _, codon := range codons {
		valFeda = append(valFeda, codon.ToFeda()...)
	}

	val, err := feda.FedaToUint(valFeda)
	if err != nil {
		log.Println("codons.Codon.CodonsToUint - Error converting codons to uint.")
	}

	return val, err
}

// UUIDToCodons return the codons that code an UUID.
func UUIDToCodons(uuid string) ([]Codon, error) {
	codons := make([]Codon, 0, len(uuid)*3)

	for _, char := range uuid {
		if char == '-' {
			codonBuilder := NewBuilder()
			codonBuilder.AddFid(feda.SPACE)
			codonBuilder.AddFid(feda.Onn)
			codonBuilder.AddFid(feda.SPACE)

			codons = append(codons, codonBuilder.GetCodon())
		} else {
			fid, err := feda.HexatoFid(char)
			if err != nil {
				return nil, err
			}
			codonBuilder := NewBuilder()
			codonBuilder.AddFid(feda.SPACE)
			codonBuilder.AddFid(fid)
			codonBuilder.AddFid(feda.SPACE)

			codons = append(codons, codonBuilder.GetCodon())
		}
	}

	return codons, nil
}

// CodonsToUUID return the UUID that is coded by the codons.
func CodonsToUUID(codons []Codon) (string, error) {
	if len(codons) == 0 {
		return "", nil
	}

	if len(codons) != 36 {
		return "", ErrInvalidUUIDCodons
	}

	var newStringBuinder strings.Builder

	for _, codon := range codons {
		if (codon[0] != feda.SPACE) || (codon[2] != feda.SPACE) {
			return "", ErrInvalidUUIDCodons
		}

		if codon[1] == feda.Onn {
			newStringBuinder.WriteString("-")
		} else {
			fid := codon[1]
			char, err := fid.ToHexa()
			if err != nil {
				return "", err
			}
			newStringBuinder.WriteString(string(char))
		}
	}

	return newStringBuinder.String(), nil
}
