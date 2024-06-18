package components

import (
	"github.com/botscubes/bot-components/context"
)

type ToIntComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Source string `json:"source"`
	} `json:"data"`
}

func (mc *ToIntComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *ToIntComponent) Execute(ctx *context.Context) (*any, error) {
	var val, err = ctx.GetValue(mc.Data.Source)
	if err != nil {
		return nil, err
	}
	i, err := val.ToInt64()
	if err != nil {
		return nil, err
	}
	var a any = i
	return &a, nil
}
