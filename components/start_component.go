package components

import (
	"github.com/botscubes/bot-components/context"
)

type StartComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
}

func (sc *StartComponent) GetOutputs() Outputs {
	return &sc.Outputs
}

func (sc *StartComponent) Execute(ctx *context.Context) error {

	return nil
}
