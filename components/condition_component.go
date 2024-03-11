package components

import "github.com/botscubes/bot-components/context"

type ConditionOutputs struct {
	ComponentOutputs

	IdIfFalse int64 `json:"idIfFalse"`
}

type ConditionComponent struct {
	ComponentData

	Outputs ConditionOutputs `json:"outputs"`
}

func (cc *ConditionComponent) GetOutputs() Outputs {
	return &cc.Outputs
}

func (cc *ConditionComponent) Execute(ctx *context.Context) error {
	v, err := ctx.GetValue(cc.Path)
	if err != nil {
		return err
	}
	b, err := v.ToBool()
	if err != nil {
		return err
	}
	if !b {
		cc.Outputs.NextComponentId = &cc.Outputs.IdIfFalse
	}
	return nil
}
