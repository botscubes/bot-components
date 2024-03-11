package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
)

type FormatComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		FormatString string `json:"formatString"`
	} `json:"data"`
}

func (fc *FormatComponent) GetOutputs() Outputs {
	return &fc.Outputs
}

func (fc *FormatComponent) Execute(ctx *context.Context) (*any, error) {
	var s any
	s, err := format.Format(fc.Data.FormatString, ctx)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
