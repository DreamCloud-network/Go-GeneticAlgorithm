package place

import (
	"log"
)

type ThingType uint

const (
	Unknown ThingType = iota
	Life
	DNA
	Light
	Earth
	Nutrient
	Energy
	TimeQuantum
	KeyboardAutomata
	KeyboardInput
	KeyboardOutput
	KeyboardEcho
	TerminalAutomata
	Supplier
	Money
)

// String returns the string representation of the thing type
func (thing ThingType) String() (string, error) {
	switch thing {
	case Unknown:
		return "Unknown", nil
	case Life:
		return "Life", nil
	case DNA:
		return "DNA", nil
	case Light:
		return "Light", nil
	case Earth:
		return "Earth", nil
	case Nutrient:
		return "Nutrient", nil
	case Energy:
		return "Energy", nil
	case TimeQuantum:
		return "TimeQuantum", nil
	case KeyboardAutomata:
		return "KeyboardAutomata", nil
	case KeyboardInput:
		return "KeyboardInput", nil
	case KeyboardOutput:
		return "KeyboardOutput", nil
	case KeyboardEcho:
		return "KeyboardEcho", nil
	case TerminalAutomata:
		return "TerminalAutomata", nil
	case Supplier:
		return "Supplier", nil
	case Money:
		return "Money", nil
	default:
		log.Println("items.String - invalid item")
		err := ErrorInvalidThingType
		return err.Error(), err
	}
}
