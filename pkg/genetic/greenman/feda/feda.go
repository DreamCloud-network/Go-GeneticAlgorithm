package feda

import (
	"math"
)

type Fid rune

const (
	SPACE Fid = iota + ' '

	// Fid => Feda
	// Aicmí 1
	Beith
	Luis
	Fearn
	Saille
	Nuin

	// Aicmí 2
	Uath
	Duir
	Tinne
	Coll
	Cert

	// Aicmí 3
	Muin
	Gort
	Getal
	Straif
	Ruis

	// Aicmí 4
	Ailm
	Onn
	Ur
	Edad
	Idad

	// Forfid => Forfeda
	Ebad
	Oir
	Uillenn
	Ifín
	Loch
	Peith // This feda represents what is lost, in therms of not finding itself. It is the end.

	INITIATOR
	REVERSE_INITIATOR
)

func (fid Fid) String() string {

	if fid == INITIATOR {
		return string(fid)
	} else {
		return /*string(SPACE)*/ string(' ') + string(fid)
	}
}

/*
type Fid byte

const (
	SPACE Fid = iota
	INITIATOR
	TERMINATOR

	// Fid => Feda
	// Aicmí 1
	Beith
	Luis
	Fearn
	Saille
	Nuin

	// Aicmí 2
	Uath
	Duir
	Tinne
	Coll
	Cert

	// Aicmí 3
	Muin
	Gort
	Getal
	Straif
	Ruis

	// Aicmí 4
	Ailm
	Onn
	Ur
	Edad
	Idad

	// Forfid => Forfeda
	Ebad
	Oir
	Uillenn
	Ifín
	Loch
	Peith // This feda represents what is lost, in therms of not finding itself. It is the end.
)

func (fid Fid) String() string {
	var fedaStr strings.Builder

	fedaStr.WriteString(" ")

	switch fid {
	case INITIATOR:
		fedaStr.Reset()
		fedaStr.WriteString("᚛")
		return fedaStr.String()
	case SPACE:
		fedaStr.Reset()
		fedaStr.WriteString(" ")
		return fedaStr.String()
	case TERMINATOR:
		fedaStr.Reset()
		fedaStr.WriteString("᚜")
		return fedaStr.String()

	case Beith:
		fedaStr.WriteString("ᚁ")
	case Luis:
		fedaStr.WriteString("ᚂ")
	case Fearn:
		fedaStr.WriteString("ᚃ")
	case Saille:
		fedaStr.WriteString("ᚄ")
	case Nuin:
		fedaStr.WriteString("ᚅ")

	case Uath:
		fedaStr.WriteString("ᚆ")
	case Duir:
		fedaStr.WriteString("ᚇ")
	case Tinne:
		fedaStr.WriteString("ᚈ")
	case Coll:
		fedaStr.WriteString("ᚉ")
	case Cert:
		fedaStr.WriteString("ᚊ")

	case Muin:
		fedaStr.WriteString("ᚋ")
	case Gort:
		fedaStr.WriteString("ᚌ")
	case Getal:
		fedaStr.WriteString("ᚍ")
	case Straif:
		fedaStr.WriteString("ᚎ")
	case Ruis:
		fedaStr.WriteString("ᚏ")

	case Ailm:
		fedaStr.WriteString("ᚐ")
	case Onn:
		fedaStr.WriteString("ᚑ")
	case Ur:
		fedaStr.WriteString("ᚒ")
	case Edad:
		fedaStr.WriteString("ᚓ")
	case Idad:
		fedaStr.WriteString("ᚔ")

	case Ebad:
		fedaStr.WriteString("ᚕ")
	case Oir:
		fedaStr.WriteString("ᚖ")
	case Uillenn:
		fedaStr.WriteString("ᚗ")
	case Ifín:
		fedaStr.WriteString("ᚘ")
	case Loch:
		fedaStr.WriteString("ᚙ")
	case Peith:
		fedaStr.WriteString("ᚚ")

	default:
		return "Unknown"
	}

	fedaStr.WriteString(" ")

	return fedaStr.String()
}
*/
// IsFeda returns true if a feda.
func (f Fid) IsFid() bool {
	return f >= Beith && f <= Idad
}

// IsForfeda returns true if the feda is a forfeda.
func (f Fid) IsForfid() bool {
	return f >= Ebad && f <= Peith
}

// IsSignalization returns true if the fid is a signalization.
func (f Fid) IsSignalization() bool {
	switch f {
	case SPACE:
	case INITIATOR:
	case REVERSE_INITIATOR:
		return true
	}
	return false
}

// IsInitiator returns true if it is a feda, forfeda or signalization.
func (f Fid) IsValid() bool {
	//return f >= SPACE && f <= Peith
	return f >= SPACE && f <= REVERSE_INITIATOR
}

// Converts an unsigned int to a feda array
// Most significat fid in the left.
func UintToFeda(val uint) []Fid {

	if val == 0 {
		return []Fid{SPACE}
	}

	fedaNum := make([]Fid, 0)

	base := uint(Idad) - uint(Beith) + 1

	for val > 0 {
		remainder := val % base
		val = val / base
		if remainder > 0 {
			newFid := Fid(remainder + uint(SPACE))
			fedaNum = append([]Fid{newFid}, fedaNum...)
		} else if val > 0 {
			fedaNum = append([]Fid{SPACE}, fedaNum...)
		}
	}

	return fedaNum
}

// Converts a fid array to a unsigned integer.
// The most significant digit is on the lefth.
func FedaToUint(feda []Fid) (uint, error) {

	if len(feda) == 0 {
		return 0, ErrEmptyArray
	}

	value := uint(0)

	pot := (uint(len(feda)) - 1)

	base := uint(Idad) - uint(Beith) + 1

	for cont := 0; cont < len(feda); cont++ {
		fidValue := uint(feda[cont]) - uint(SPACE)
		value += fidValue * uint(math.Pow(float64(base), float64(pot)))

		pot--
	}

	return value, nil
}

/*
// Converts an unsigned int to a feda array
// Most significat fid in the left.
func UintToFeda(val uint) []Fid {

	if val == 0 {
		return []Fid{Beith}
	}

	fedaNum := make([]Fid, 0)

	base := uint(Idad) - uint(Beith) + 1

	for val > 0 {
		remainder := val % base
		val = val / base
		if remainder > 0 {
			newFid := Fid(remainder + uint(Beith))
			fedaNum = append([]Fid{newFid}, fedaNum...)
		} else if val > 0 {
			fedaNum = append([]Fid{Beith}, fedaNum...)
		}
	}

	return fedaNum
}

// Converts a fid array to a unsigned integer.
// The most significant digit is on the lefth.
func FedaToUint(feda []Fid) (uint, error) {

	if len(feda) == 0 {
		return 0, ErrEmptyArray
	}

	value := uint(0)

	pot := (uint(len(feda)) - 1)

	base := uint(Idad) - uint(Beith) + 1

	for cont := 0; cont < len(feda); cont++ {
		fidValue := uint(feda[cont]) - uint(Beith)
		value += fidValue * uint(math.Pow(float64(base), float64(pot)))

		pot--
	}

	return value, nil
}
*/
