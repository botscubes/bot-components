package components

import (
	"encoding/json"

	"github.com/botscubes/bot-components/context"
	"github.com/botscubes/bot-components/io"
)

type Component interface {
	GetNextComponentId() *int64
	GetIdIfError() *int64
	GetPath() string
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

type ComponentData struct {
	ComponentTypeData

	NextComponentId *int64 `json:"nextComponentId"`
	IdIfError       *int64 `json:"idIfError"`
	Path            string `json:"path"`
}

func (cd *ComponentData) GetNextComponentId() *int64 {
	return cd.NextComponentId
}

func (cd *ComponentData) GetPath() string {
	return cd.Path
}

func (cd *ComponentData) GetIdIfError() *int64 {
	return cd.IdIfError
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
	default:
		return nil, ErrComponentTypeNotExist
	}
}
