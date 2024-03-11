package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
	"github.com/botscubes/bot-components/io"
)

type MessageComponent struct {
	ComponentData

	Text string `json:"text"`
}

func (fc *MessageComponent) Execute(ctx *context.Context, io io.IO) error {
	var s string
	s, err := format.Format(fc.Text, ctx)
	if err != nil {
		return err
	}
	io.OutputText(s)
	return nil
}
