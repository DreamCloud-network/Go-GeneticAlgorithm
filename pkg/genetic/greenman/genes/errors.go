package genes

import "errors"

var (
	ErrNotPermited     = errors.New("not permited")
	ErrUnexpectedCodon = errors.New("unexpected codon")
	ErrUnexpectedError = errors.New("unexpected error")
)
