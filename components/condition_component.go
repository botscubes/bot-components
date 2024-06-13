package components

import (
	"strings"

	"github.com/botscubes/bot-components/context"
)

type ConditionOutputs struct {
	ComponentOutputs

	IdIfFalse int64 `json:"idIfFalse"`
}

type ConditionComponent struct {
	ComponentData
	Data struct {
		Expression string `json:"expression"`
	} `json:"data"`
	Outputs ConditionOutputs `json:"outputs"`
}

func (cc *ConditionComponent) GetOutputs() Outputs {
	return &cc.Outputs
}

func (cc *ConditionComponent) Execute(ctx *context.Context) error {
	expression := strings.TrimSpace(cc.Data.Expression)
	if expression == "true" {
		return nil
	}
	if expression == "false" {
		cc.Outputs.NextComponentId = &cc.Outputs.IdIfFalse
	}
	value, err := ctx.GetValue(expression)
	if err != nil {
		return err
	}
	str, err := value.ToString()
	if err != nil {
		return err
	}
	if str != "true" {
		cc.Outputs.NextComponentId = &cc.Outputs.IdIfFalse
	}
	return nil
}
