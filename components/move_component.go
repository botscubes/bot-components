package components

import (
	"github.com/botscubes/bot-components/context"
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
func (mc *MoveComponent) Execute(ctx *context.Context) (*any, error) {
	var val, err = ctx.GetValue(mc.Data.Source)
	if err != nil {
		return nil, err
	}
	var a any = val.GetRawValue()
	return &a, nil
}
