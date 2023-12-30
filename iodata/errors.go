package iodata

import (
	"errors"
	"fmt"
)

var (
	ErrTypeAssertion      = errors.New("Type assertion")
	ErrTypeNotFound       = errors.New("Type not found")
	ErrIndexOutOfRange    = errors.New("Index out of range")
	ErrNoPropertyInObject = errors.New("No property in object")
)

var (
	ErrVariableNameBeginning    = errors.New("The variable name must begin with a letter or underscore")
	ErrVariableName             = errors.New("The variable name must contain only letters, numbers or underscores")
	ErrNoClosingSquareBracket   = errors.New("No closing square bracket")
	ErrWrongIndex               = errors.New("Wrong index")
	ErrUnknownCharacter         = errors.New("Unknown character")
	ErrVariableNameNotSpecified = errors.New("Variable name not specified")
)

func NewErrTypeAssertion(currentType string, expextedType string) error {
	return fmt.Errorf("%w: Current type: %s, Expexted type: %s", ErrTypeAssertion, currentType, expextedType)
}
func NewErrIndexOutOfRange(arrayName string, currentIndex int, from int, to int) error {
	return fmt.Errorf("%w: array name: %s, index: %d, range: [%d:%d]",
		ErrIndexOutOfRange,
		arrayName,
		currentIndex,
		from,
		to)
}
func NewErrNoPropertyInObject(obj string, property string) error {
	return fmt.Errorf("%w: object name: %s, property: %s", ErrNoPropertyInObject, obj, property)
}
