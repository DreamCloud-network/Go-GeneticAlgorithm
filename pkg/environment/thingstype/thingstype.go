package thingstype

import (
	"errors"
	"log"
)

var (
	ErrorInvalidThingType = errors.New("invalid thing type")
)

type ThingType int

const (
	Landscape ThingType = iota
	Robot
	Cell
	Trash
	Food
	Unknown
)

// String returns the string representation of the thing type
func (m ThingType) String() (string, error) {
	switch m {
	case Landscape:
		return "Landscape", nil
	case Robot:
		return "Robot", nil
	case Trash:
		return "Trash", nil
	case Unknown:
		return "Unknown", nil
	default:
		log.Println("items.String - invalid item")
		err := ErrorInvalidThingType
		return err.Error(), err
	}
}
