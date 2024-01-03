package components

import (
	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
)

type FormatComponent struct {
	ComponentData

	FormatString string `json:"formatString"`
}

func (fc *FormatComponent) Execute(ctx *context.Context) (*any, error) {
	var s any
	s, err := format.Format(fc.FormatString, ctx)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
