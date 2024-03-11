package components

import (
	"errors"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type TextInputComponent struct {
	ComponentData

	Text string `json:"text"`
}

func (fc *TextInputComponent) Execute(ctx *context.Context, io io.IO) (*any, error) {
	s := io.InputText()
	if s == nil {
		return nil, errors.New("No text was entered")
	}
	var a any
	a = *s
	return &a, nil
}
