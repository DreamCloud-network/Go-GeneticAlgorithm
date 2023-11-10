package dnastrand

import "errors"

var (
	ErrDifferentGeneCount     = errors.New("different gene count")
	ErrInvalidDNAStrandFormat = errors.New("invalid DNA Strand format")
)
