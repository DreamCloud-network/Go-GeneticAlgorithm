package codons

import "errors"

var (
	ErrorInvalidFid       = errors.New("invalid fid")
	ErrorInvalidForfid    = errors.New("invalid forfid")
	ErrorCodonFull        = errors.New("codon full")
	ErrorCodonNotComplete = errors.New("codon not complete")
	ErrorInvalidState     = errors.New("invalid state")
	ErrorInvalidAction    = errors.New("invalid action")
)
