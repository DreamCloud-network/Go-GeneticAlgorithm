package basicmachinery

import "errors"

var (
	ErrInvalidGeneCode = errors.New("invalid gene code")
	ErrNoExternalEnv   = errors.New("no external environment")
	ErrNoInternalEnv   = errors.New("no internal environment")
)
