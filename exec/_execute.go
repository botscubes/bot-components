package exec

import (
	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
)

func Execute(ctx *context.Context, component components.Component) (*int, error) {
	switch cmp := component.(type) {
	case components.ActionComponent:

		v, err := cmp.Execute(ctx)
		if err != nil {
			return nil, err
		}
		ctx.SetValue(component.GetPath(), *v)
	case components.ControlComponent:
		err := cmp.ChangeNextComponentId(ctx)
		if err != nil {
			return nil, err
		}

	//var component components.Component
	//switch tp {
	//case components.TypeFormat:
	//	var format components.FormatComponent
	//	err := json.Unmarshal(jsonData, &format)
	//	if err != nil {
	//		return 0, err
	//	}
	//	component = &format
	default:
		return nil, ErrComponentNotImplInterface
	}

	nextId := component.GetNextComponentId()

	return nextId, nil
}
