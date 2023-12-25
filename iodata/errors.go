package iodata

import "errors"

var (
	ErrTypeMismatch    = errors.New("Type mismatch")
	ErrTypeNotFound    = errors.New("Type not found")
	ErrIndexOutOfRange = errors.New("Index out of range")
	ErrKeyIsNotInMap   = errors.New("The key is not in the map")
)
