package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
	"github.com/botscubes/bot-components/io"
)

type PhotoComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
}

func (mc *PhotoComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *PhotoComponent) Execute(ctx *context.Context, io io.IO) error {
	var s string
	s, err := format.Format(mc.Path, ctx)
	if err != nil {
		return err
	}
	err = io.PrintPhoto([]byte(s))
	return err
}
