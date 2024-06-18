package components

import (
	"errors"
	"sort"
	"strconv"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type ButtonComponentOutputs struct {
	outputs map[string]*int64
}

type ButtonComponent struct {
	ComponentData

	Outputs map[string]*int64 `json:"outputs"`

	Data struct {
		Text    string                    `json:"text"`
		Buttons map[string]*io.ButtonData `json:"buttons"`
	} `json:"data"`
}

func (o *ButtonComponentOutputs) GetNextComponentId() *int64 {
	id, ok := o.outputs["nextComponentId"]
	if ok {
		return id
	}
	return nil
}
func (o *ButtonComponentOutputs) GetIdIfError() *int64 {
	id, ok := o.outputs["idIfError"]
	if ok {
		return id
	}
	return nil
}
func (bc *ButtonComponent) GetOutputs() Outputs {
	return &ButtonComponentOutputs{
		outputs: bc.Outputs,
	}
}

func (bc *ButtonComponent) Execute(ctx *context.Context, inout io.IO) (*any, error) {
	text := inout.ReadText()
	if text == nil {
		bc.Outputs["nextComponentId"] = bc.Id
		keys := make([]int, 0, len(bc.Data.Buttons))
		for key := range bc.Data.Buttons {
			numKey, err := strconv.Atoi(key)
			if err != nil {
				return nil, errors.New("Error converting key to int: " + err.Error())
			}
			keys = append(keys, numKey)
		}
		sort.Ints(keys)
		buttons := make([]*io.ButtonData, len(keys))
		for i, key := range keys {
			buttons[i] = bc.Data.Buttons[strconv.Itoa(key)]
		}

		err := inout.PrintButtons(bc.Data.Text, buttons)
		return nil, err
	}
	if *text == "" {
		return nil, errors.New("Empty string entered")
	}
	buttons := make(map[string]*int64)
	for key, button := range bc.Data.Buttons {
		next, ok := bc.Outputs[key]
		if ok {
			buttons[button.Text] = next
		}
	}
	next, ok := buttons[*text]
	if !ok {
		return nil, errors.New("Button not selected")
	}
	bc.Outputs["nextComponentId"] = next
	var a any
	a = *text
	return &a, nil
}
