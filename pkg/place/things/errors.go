package things

import "errors"

var (
	ErrorPlaceNotNil = errors.New("place is not nil. Must use place.MoveThingToPlace")
)
