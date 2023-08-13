package codons

import "errors"

var (
	ErrorInvalidFeda      = errors.New("invalid feda")
	ErrorCodonFull        = errors.New("codon full")
	ErrorCodonNotComplete = errors.New("codon not complete")
	ErrorInvalidState     = errors.New("invalid state")
)
