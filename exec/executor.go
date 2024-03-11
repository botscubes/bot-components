package exec

import (
	"github.com/botscubes/bot-components/components"
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type Executor struct {
	io  io.IO
	ctx *context.Context
}

func NewExecutor(ctx *context.Context, io io.IO) *Executor {
	return &Executor{
		io,
		ctx,
	}
}

func (e *Executor) Execute(component components.Component) (*int64, error) {
	switch cmp := component.(type) {
	case components.ActionComponent:

		v, err := cmp.Execute(e.ctx)
		if err != nil {
			return cmp.GetOutputs().GetIdIfError(), err
		}
		e.ctx.SetValue(component.GetPath(), *v)
	case components.ControlComponent:
		err := cmp.Execute(e.ctx)
		if err != nil {
			return cmp.GetOutputs().GetIdIfError(), err
		}
	case components.InputComponent:
		v, err := cmp.Execute(e.ctx, e.io)
		if err != nil {
			return cmp.GetOutputs().GetIdIfError(), err
		}
		e.ctx.SetValue(component.GetPath(), *v)
	case components.OutputComponent:
		err := cmp.Execute(e.ctx, e.io)
		if err != nil {
			return cmp.GetOutputs().GetIdIfError(), err
		}

	default:
		return nil, ErrComponentNotImplInterface
	}

	nextId := component.GetOutputs().GetNextComponentId()

	return nextId, nil
}
