package circuit

import "errors"

var (
	ErrOutOfRange = errors.New("out of range")
	ErrSameNode   = errors.New("not allowed to connect the node with himself")
)
