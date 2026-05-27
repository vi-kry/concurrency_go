package compute

import "errors"

var (
	errEmptyQuery     = errors.New("empty query")
	errInvalidCommand = errors.New("invalid command")
	errInvalidArgs    = errors.New("invalid arguments")
)
