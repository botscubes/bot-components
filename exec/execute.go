package exec

import (
	"encoding/json"

	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
)

func Execute(ctx *context.Context, tp components.ComponentType, jsonData []byte) (int, error) {
	var component components.Component
	switch tp {
	case components.TypeFormat:
		var format components.FormatComponent
		err := json.Unmarshal(jsonData, &format)
		if err != nil {
			return 0, err
		}
		component = &format
	default:
		return 0, NewErrComponentTypeNotExist(tp)
	}

	//savePath := component.GetSavePath()

	nextId := component.GetNextComponentId()

	return nextId, nil
}
