package components

import "github.com/botscubes/bot-components/context"

type ConditionComponent struct {
	ComponentData

	IdIfFalse int    `json:"idIfFalse"`
	CheckPath string `json:"checkPath"`
}

func (cc *ConditionComponent) ChangeNextComponentId(ctx *context.Context) error {
	v, err := ctx.GetValue(cc.CheckPath)
	if err != nil {
		return err
	}
	b, err := v.ToBool()
	if err != nil {
		return err
	}
	if !b {
		cc.NextComponentId = cc.IdIfFalse
	}
	return nil
}
