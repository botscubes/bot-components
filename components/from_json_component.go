package components

import (
	"encoding/json"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/format"
)

type FromJSONComponent struct {
	ComponentData

	Outputs ComponentOutputs `json:"outputs"`
	Data    struct {
		Json string `json:"json"`
	} `json:"data"`
}

func (c *FromJSONComponent) GetOutputs() Outputs {
	return &c.Outputs
}
func (c *FromJSONComponent) Execute(ctx *context.Context) (*any, error) {
	json_str, err := format.Format(c.Data.Json, ctx)
	if err != nil {
		return nil, err
	}
	var m map[string]any
	err = json.Unmarshal([]byte(json_str), &m)
	if err != nil {
		return nil, err
	}
	var a any = m
	return &a, nil
}
