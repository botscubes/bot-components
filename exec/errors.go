package exec

import (
	"errors"
	"fmt"

	"github.com/botscubes/bot-components/components"
)

var (
	ErrComponenTypeNotExist = errors.New("Component type does not exist")
)

func NewErrComponentTypeNotExist(tp components.ComponentType) error {
	return fmt.Errorf("%w: %s", ErrComponenTypeNotExist, tp)
}
