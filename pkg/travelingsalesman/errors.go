package travelingsalesman

import "errors"

var (
	ErrNotAlive        = errors.New("traveling salesman is not alive")
	ErrInvalidPosition = errors.New("invalid position")
	ErrInvalidMove     = errors.New("invalid move")
)
