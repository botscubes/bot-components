package exec

import (
	"errors"
)

var (
	ErrComponentNotImplInterface = errors.New("The component does not implement the interface")
)
