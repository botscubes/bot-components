package components

import (
	"encoding/json"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type Component interface {
	GetPath() string
	GetOutputs() Outputs
}
type Outputs interface {
	GetNextComponentId() *int64
	GetIdIfError() *int64
}

type (
	ActionComponent interface {
		Component

		Execute(ctx *context.Context) (*any, error)
	}

	ControlComponent interface {
		Component

		Execute(ctx *context.Context) error
	}

	InputComponent interface {
		Component

		Execute(ctx *context.Context, io io.IO) (*any, error)
	}
	OutputComponent interface {
		Component

		Execute(ctx *context.Context, io io.IO) error
	}
)

type ComponentTypeData struct {
	Type ComponentType `json:"type"`
}

type ComponentOutputs struct {
	NextComponentId *int64 `json:"nextComponentId"`
	IdIfError       *int64 `json:"idIfError"`
}

func (co *ComponentOutputs) GetNextComponentId() *int64 {
	return co.NextComponentId
}
func (co *ComponentOutputs) GetIdIfError() *int64 {
	return co.IdIfError
}

type ComponentData struct {
	ComponentTypeData
	Path string `json:"path"`
}

func (cd *ComponentData) GetPath() string {
	return cd.Path
}

func NewComponentFromJSON(tp ComponentType, jsonData []byte) (Component, error) {
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

	case TypeMessage:
		var m MessageComponent
		err := json.Unmarshal(jsonData, &m)
		if err != nil {
			return nil, err
		}
		return &m, err

	default:
		return nil, ErrComponentTypeNotExist
	}
}
