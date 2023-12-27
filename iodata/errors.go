package iodata

import "errors"

var (
	ErrTypeMismatch    = errors.New("Type mismatch")
	ErrTypeNotFound    = errors.New("Type not found")
	ErrIndexOutOfRange = errors.New("Index out of range")
	ErrKeyIsNotInMap   = errors.New("The key is not in the map")
)

var (
	ErrVariableNameBeginning    = errors.New("The variable name must begin with a letter or underscore")
	ErrVariableName             = errors.New("The variable name must contain only letters, numbers or underscores")
	ErrNoClosingSquareBracket   = errors.New("No closing square bracket")
	ErrWrongIndex               = errors.New("Wrong index")
	ErrUnknownCharacter         = errors.New("Unknown character")
	ErrVariableNameNotSpecified = errors.New("Variable name not specified")
)
