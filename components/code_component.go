package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
	"github.com/botscubes/bql/api"
)

type CodeComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Code string `json:"code"`
	} `json:"data"`
}

func (mc *CodeComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *CodeComponent) Execute(ctx *context.Context, io io.IO) (*any, error) {
	keys := ctx.GetKyes()
	v, err := api.EvalWithCtx(mc.Data.Code, ctx, &keys)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
