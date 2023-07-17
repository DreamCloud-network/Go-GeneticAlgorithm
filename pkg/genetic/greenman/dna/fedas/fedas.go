package fedas

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
	INI Feda = iota

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
	Ebad Feda = iota
	Oir
	Uillenn
	Ifín
	Loch
	Peith // This feda represents what is lost, in therms of not finding itself. It is the end.
)

func (feda Feda) String() string {
	if feda == INI {
		return "᚛"
	}

	str := " "
	switch feda {
	case Beith:
		str += "ᚁ"
	case Luis:
		str += "ᚂ"
	case Fearn:
		str += "ᚃ"
	case Saille:
		str += "ᚄ"
	case Nuin:
		str += "ᚅ"
	case Uath:
		str += "ᚆ"
	case Duir:
		str += "ᚇ"
	case Tinne:
		str += "ᚈ"
	case Coll:
		str += "ᚉ"
	case Cert:
		str += "ᚊ"
	case Muin:
		str += "ᚋ"
	case Gort:
		str += "ᚌ"
	case Getal:
		str += "ᚍ"
	case Straif:
		str += "ᚎ"
	case Ruis:
		str += "ᚏ"
	case Ailm:
		str += "ᚐ"
	case Onn:
		str += "ᚑ"
	case Ur:
		str += "ᚒ"
	case Edad:
		str += "ᚓ"
	case Idad:
		str += "ᚔ"
	case Ebad:
		str += "ᚕ"
	case Oir:
		str += "ᚖ"
	case Uillenn:
		str += "ᚗ"
	case Ifín:
		str += "ᚘ"
	case Loch:
		str += "ᚙ"
	case Peith:
		str += "ᚚ"
	default:
		return "Unknown"
	}

	str += " "

	return str
}

/*
type Aicme1 byte

const (
	Beith Aicme1 = iota + 1
	Luis
	Fearn
	Saille
	Nuin
)

func (aicme Aicme1) String() string {
	switch aicme {
	case Beith:
		return "-ᚁ-"
	case Luis:
		return "-ᚂ-"
	case Fearn:
		return "-ᚃ-"
	case Saille:
		return "-ᚄ-"
	case Nuin:
		return "-ᚅ-"
	default:
		return "Unknown"
	}
}

type Aicme2 byte

const (
	Uath Aicme2 = iota + 1
	Duir
	Tinne
	Coll
	Cert
)

func (aicme Aicme2) String() string {
	switch aicme {
	case Uath:
		return "-ᚆ-"
	case Duir:
		return "-ᚇ-"
	case Tinne:
		return "-ᚈ-"
	case Coll:
		return "-ᚉ-"
	case Cert:
		return "-ᚊ-"
	default:
		return "Unknown"
	}
}

type Aicme3 byte

const (
	Muin Aicme3 = iota + 1
	Gort
	Getal
	Straif
	Ruis
)

func (aicme Aicme3) String() string {
	switch aicme {
	case Muin:
		return "-ᚋ-"
	case Gort:
		return "-ᚌ-"
	case Getal:
		return "-ᚍ-"
	case Straif:
		return "-ᚎ-"
	case Ruis:
		return "-ᚏ-"
	default:
		return "Unknown"
	}
}

type Aicme4 byte

const (
	Ailm Aicme4 = iota + 1
	Onn
	Ur
	Edad
	Idad
)

func (aicme Aicme4) String() string {
	switch aicme {
	case Ailm:
		return "-ᚐ-"
	case Onn:
		return "-ᚑ-"
	case Ur:
		return "-ᚒ-"
	case Edad:
		return "-ᚓ-"
	case Idad:
		return "-ᚔ-"
	default:
		return "Unknown"
	}
}

type Forfeda byte

const (
	Ebad Forfeda = iota
	Oir
	Uillenn
	Ifín
	Loch
	Peith
)

func (aicme Forfeda) String() string {
	switch aicme {
	case Ebad:
		return "-ᚕ-"
	case Oir:
		return "-ᚖ-"
	case Uillenn:
		return "-ᚗ-"
	case Ifín:
		return "-ᚘ-"
	case Loch:
		return "-ᚙ-"
	case Peith:
		return "-ᚚ-"
	default:
		return "Unknown"
	}
}
*/
