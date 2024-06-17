package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type MoveComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Source string `json:"source"`
	} `json:"data"`
}

func (mc *MoveComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *MoveComponent) Execute(ctx *context.Context, io io.IO) (*any, error) {
	var val, err = ctx.GetValue(mc.Data.Source)
	if err != nil {
		return nil, err
	}
	var a any = val.GetRawValue()
	return &a, nil
}
