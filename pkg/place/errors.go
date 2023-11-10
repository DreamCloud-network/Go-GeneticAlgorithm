package place

import "errors"

var (
	ErrorInvalidThingType       = errors.New("invalid thing type")
	ErrorNoConnection           = errors.New("no connection")
	ErrorInvalidConnectionIndex = errors.New("invalid connection index")
	ErrorThingNotInPlace        = errors.New("thing is not in place")
	ErrorThingIsNil             = errors.New("thing is nil")
)
