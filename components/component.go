package components

import (
	"encoding/json"

	"github.com/botscubes/bot-components/context"
)

type Component interface {
	GetNextComponentId() *int
	GetPath() string
}

type (
	ActionComponent interface {
		Component

		Execute(ctx *context.Context) (*any, error)
	}

	ControlComponent interface {
		Component

		ChangeNextComponentId(ctx *context.Context) error
	}

	InputComponent interface {
		Component

		Input(ctx *context.Context) (*any, error)
	}

	OutputComponent interface {
		Component

		Output(ctx *context.Context) error
	}
)

type ComponentTypeData struct {
	Type ComponentType `json:"componentType"`
}

type ComponentData struct {
	ComponentTypeData

	NextComponentId *int   `json:"nextComponentId"`
	Path            string `json:"path"`
}

func (cd *ComponentData) GetNextComponentId() *int {
	return cd.NextComponentId
}

func (cd *ComponentData) GetPath() string {
	return cd.Path
}

func NewActionOrControlComponentFromJSON(tp ComponentType, jsonData []byte) (Component, error) {
	switch tp {
	case TypeFormat:
		var f FormatComponent
		err := json.Unmarshal(jsonData, &f)
		if err != nil {
			return nil, err
		}
		return &f, nil
	case TypeCondition:
		var c ConditionComponent
		err := json.Unmarshal(jsonData, &c)
		if err != nil {
			return nil, err
		}
		return &c, err
	default:
		return nil, ErrComponentTypeNotExist
	}
}
