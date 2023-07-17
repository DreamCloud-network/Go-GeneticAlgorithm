package robby

type Action byte

const (
	StepSouth Action = iota
	StepNorth
	StepEast
	StepWest
	RandomMove
	Pickup
	DoNothing
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
