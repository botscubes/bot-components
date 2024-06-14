package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type ToIntComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Source      string `json:"source"`
		Destination string `json:"destination"`
	} `json:"data"`
}

func (mc *ToIntComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *ToIntComponent) Execute(ctx *context.Context, io io.IO) (*any, error) {
	var val, err = ctx.GetValue(mc.Data.Source)
	if err != nil {
		return nil, err
	}
	i, err := val.ToInt64()
	if err != nil {
		return nil, err
	}
	mc.Path = mc.Data.Destination
	var a any = i
	return &a, nil
}
