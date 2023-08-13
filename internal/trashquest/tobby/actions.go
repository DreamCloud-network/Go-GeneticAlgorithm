package tobby

type Action byte

const (
	RandomMove Action = iota + 1
	StepNorth
	StepWest
	StepSouth
	StepEast
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
	case Pickup:
		return "Pickup"

	default:
		err := ErrorInvalidAction
		return err.Error()
	}
}
