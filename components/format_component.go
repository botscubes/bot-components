package components

import "github.com/botscubes/bot-components/context"

type FormatComponent struct {
	FormatString string `json:"formatString"`
	ComponentData
}

func (fc *FormatComponent) Execute(ctx *context.Context) (any, error) {

	return nil, nil
}
