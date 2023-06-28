package trashquest

import (
	"errors"
	"log"
)

type Action int

const (
	StepNorth Action = iota
	StepSouth
	StepEast
	StepWest
	RandomMove
	DoNothing
	Pickup
)

var (
	ErrorInvalidAction = errors.New("invalid action")
)

func (m Action) String() (string, error) {
	switch m {
	case StepNorth:
		return "StepNorth", nil
	case StepSouth:
		return "StepSouth", nil
	case StepEast:
		return "StepEast", nil
	case StepWest:
		return "StepWest", nil
	case RandomMove:
		return "RandomMove", nil
	case DoNothing:
		return "DoNothing", nil
	case Pickup:
		return "Pickup", nil

	default:
		log.Println("movement.String - invalid action")
		err := ErrorInvalidAction
		return err.Error(), err
	}
}
