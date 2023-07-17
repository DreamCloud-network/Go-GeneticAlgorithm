package runes

type Rune byte

const (
	Blank Rune = iota
	Fehu
	Uruz
	Turisaz
	Ansuz
	Raido
	Kano
	Gebo
	Wunjo

	Hagalaz
	Naudiz
	Isaz
	Jera
	Thurisaz
	Phedor
	Algiz
	Souwolo

	Tiwaz
	Berkano
	Ehwaz
	Mannaz
	Laguz
	Ingwaz
	Dagaz
	Othala
)

func (rune Rune) String() string {
	str := " "
	switch rune {
	case Blank:
		str += "ᛰ "
	case Fehu:
		str += "ᚠ"
	case Uruz:
		str += "ᚢ"
	case Turisaz:
		str += "ᚦ"
	case Ansuz:
		str += "ᚨ"
	case Raido:
		str += "ᚱ"
	case Kano:
		str += "ᚲ"
	case Gebo:
		str += "ᚷ"
	case Wunjo:
		str += "ᚹ"

	case Hagalaz:
		str += "ᚺ"
	case Naudiz:
		str += "ᚾ"
	case Isaz:
		str += "ᛁ"
	case Jera:
		str += "ᛃ"
	case Thurisaz:
		str += "ᛇ"
	case Phedor:
		str += "ᛈ"
	case Algiz:
		str += "ᛉ"
	case Souwolo:
		str += "ᛊ"

	case Tiwaz:
		str += "ᛏ"
	case Berkano:
		str += "ᛒ"
	case Ehwaz:
		str += "ᛖ"
	case Mannaz:
		str += "ᛗ"
	case Laguz:
		str += "ᛚ"
	case Ingwaz:
		str += "ᛜ"
	case Dagaz:
		str += "ᛞ"
	case Othala:
		str += "ᛟ"
	default:
		str += "Unknown"
	}
	str += " "

	return str
}
