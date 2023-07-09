package genome

import "errors"

var (
	ErrActionNotAllowed = errors.New("action not allowed")
	ErrorInvalidAction  = errors.New("invalid action")
	ErrorActionsEmpty   = errors.New("actions list empty")
	ErrorGenesUndefined = errors.New("genes undefined")
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

func (m Action) String() string {
	switch m {
	case StepNorth:
		return "StepNorth"
	case StepSouth:
		return "StepSouth"
	case StepEast:
		return "StepEast"
	case StepWest:
		return "StepWest"
	case RandomMove:
		return "RandomMove"
	case DoNothing:
		return "DoNothing"
	case Pickup:
		return "Pickup"

	default:
		err := ErrorInvalidAction
		return err.Error()
	}
}

type Genes interface {
	GetActions() []Action
	Mate(genesPartner Genes) []Genes
	Duplicate() Genes
	Sequence() string
}
