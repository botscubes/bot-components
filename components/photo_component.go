package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
	"github.com/botscubes/bot-components/io"
)

type PhotoComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Name string `json:"name"`
	} `json:"data"`
}

func (mc *PhotoComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *PhotoComponent) Execute(ctx *context.Context, io io.IO) error {
	v, err := ctx.GetValue(mc.Path)
	if err != nil {
		return err
	}
	b, err := v.ToBytes()
	if err != nil {
		return err
	}
	name, err := format.Format(mc.Data.Name, ctx)
	if err != nil {
		return err
	}
	err = io.PrintPhoto(b, name)
	return err
}
