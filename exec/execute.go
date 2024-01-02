package exec

import (
	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
)

func Execute(ctx *context.Context, tp components.ComponentType, jsonData []byte) error {
	switch tp {
	case components.TypeFormat:

	default:
		return NewErrComponentTypeNotExist(tp)
	}

	return nil
}
