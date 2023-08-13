package fedas

import "strings"

/*
type Feda rune

const Ebad Feda = 'ᚕ'
const Oir Feda = 'ᚖ'
const Uillenn Feda = 'ᚗ'
const Ifín Feda = 'ᚘ'
const Loch Feda = 'ᚙ'
const Peith Feda = 'ᚚ'
const Beith Feda = 'ᚁ'
const Luis Feda = 'ᚂ'
const Fearn Feda = 'ᚃ'
const Saille Feda = 'ᚄ'
const Nuin Feda = 'ᚅ'
const Uath Feda = 'ᚆ'
const Duir Feda = 'ᚇ'
const Tinne Feda = 'ᚈ'
const Coll Feda = 'ᚉ'
const Cert Feda = 'ᚊ'
const Muin Feda = 'ᚋ'
const Gort Feda = 'ᚌ'
const Getal Feda = 'ᚍ'
const Straif Feda = 'ᚎ'
const Ruis Feda = 'ᚏ'
const Ailm Feda = 'ᚐ'
const Onn Feda = 'ᚑ'
const Ur Feda = 'ᚒ'
const Edad Feda = 'ᚓ'
const Idad Feda = 'ᚔ'
*/

type Feda byte

const (
	SPACE Feda = iota
	INITIATOR
	TERMINATOR

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

	// Forfedha
	Ebad
	Oir
	Uillenn
	Ifín
	Loch
	Peith // This feda represents what is lost, in therms of not finding itself. It is the end.
)

func (feda Feda) String() string {
	var fedaStr strings.Builder

	fedaStr.WriteString(" ")

	switch feda {
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

// IsInfoFeda returns true if a feda.
func (f Feda) IsFeda() bool {
	return f >= Beith && f <= Idad
}

// IsForfeda returns true if the feda is a forfeda.
func (f Feda) IsForfeda() bool {
	return f >= Ebad && f <= Peith
}

// IsSignalization returns true if the feda is a signalization.
func (f Feda) IsSignalization() bool {
	return f >= SPACE && f <= TERMINATOR
}

func (f Feda) IsValid() bool {
	return f >= SPACE && f <= Peith
}
