package format

import "errors"

var (
	ErrNoClosingCurlyBrace = errors.New("No closing curly brace")
	ErrNoOpeningCurlyBrace = errors.New("No opening curly brace")
)
