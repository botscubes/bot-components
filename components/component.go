package components

import "github.com/botscubes/bot-components/context"

type Component interface {
	GetNextComponentId() int
	GetSavePath() string
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

type ComponentData struct {
	Type            ComponentType `json:"componentType"`
	NextComponentId int           `json:"nextComponentId"`
	SavePath        string        `json:"savePath"`
}

func (cd *ComponentData) GetNextComponentId() int {
	return cd.NextComponentId
}

func (cd *ComponentData) GetSavePath() string {
	return cd.SavePath
}
