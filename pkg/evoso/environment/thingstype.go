package environment

import (
	"log"
)

type ThingType uint

const (
	Unknown ThingType = iota

	TimeQuantum

	Ribossome
	Collector

	Channel
	RequestSignal

	RRNA

	KeyboardAutomata
	KeyboardEchoTerminalAutomata
	TerminalAutomata

	KeyboardInput
	TerminalOutput
)

// String returns the string representation of the thing type
func (thing ThingType) String() string {
	switch thing {
	case Unknown:
		return "Unknown"

	case TimeQuantum:
		return "TimeQuantum"

	case Ribossome:
		return "Ribossome"
	case Collector:
		return "Collector"

	case RRNA:
		return "RRNA"

	case KeyboardAutomata:
		return "KeyboardAutomata"
	case KeyboardEchoTerminalAutomata:
		return "KeyboardEchoTerminalAutomata"
	case TerminalAutomata:
		return "TerminalAutomata"

	case KeyboardInput:
		return "KeyboardInput"
	case TerminalOutput:
		return "TerminalOutput"

	default:
		log.Println("items.String - invalid item")
		return "Invalid"
	}
}
