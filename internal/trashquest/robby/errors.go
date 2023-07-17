package robby

import "errors"

var (
	ErrorStrandsEmpty      = errors.New("dna strands empty")
	ErrorStrandsInvalid    = errors.New("dna strands invalid")
	ErrActionNotAllowed    = errors.New("action not allowed")
	ErrorInvalidAction     = errors.New("invalid action")
	ErrorActionsEmpty      = errors.New("actions list empty")
	ErrorGenesUndefined    = errors.New("genes undefined")
	ErrNoTrash             = errors.New("no trash in position")
	ErrorBoardUndefined    = errors.New("board undefined")
	ErrorPositionUndefined = errors.New("position undefined")
)
