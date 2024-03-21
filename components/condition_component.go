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

	if !parseExpression(cc.Data.Expression) {
		cc.Outputs.NextComponentId = &cc.Outputs.IdIfFalse
	}
	return nil
}

func parseExpression(expression string) bool {
	expression = strings.TrimSpace(expression)
	if expression == "true" {
		return true
	}

	return false
}
