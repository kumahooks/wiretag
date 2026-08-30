package wiretag

import "errors"

var (
	ErrOpen       = errors.New("could not open file")
	ErrInvalid    = errors.New("invalid audio file")
	ErrFileClosed = errors.New("audio file is closed")
)
