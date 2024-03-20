package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
	"github.com/botscubes/bot-components/io"
)

type MessageComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Text string `json:"text"`
	} `json:"data"`
}

func (mc *MessageComponent) GetOutputs() Outputs {
	return &mc.Outputs
}
func (mc *MessageComponent) Execute(ctx *context.Context, io io.IO) error {
	var s string
	s, err := format.Format(mc.Data.Text, ctx)
	if err != nil {
		return err
	}
	io.OutputText(s)
	return nil
}
