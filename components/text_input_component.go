package components

import (
	"errors"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type TextInputComponent struct {
	ComponentData
	Outputs ComponentOutputs `json:"outputs"`
}

func (tc *TextInputComponent) GetOutputs() Outputs {
	return &tc.Outputs
}

func (tc *TextInputComponent) Execute(ctx *context.Context, io io.IO) (*any, error) {
	s := io.InputText()
	if s == nil {
		tc.Outputs.NextComponentId = tc.Id
		return nil, nil
	}
	if *s == "" {
		return nil, errors.New("Empty string entered")
	}
	var a any
	a = *s
	return &a, nil
}
