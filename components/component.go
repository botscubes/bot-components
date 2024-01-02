package components

import "github.com/botscubes/bot-components/context"

type Component interface {
	GetNextComponentId() int
	GetSavePath() string
	Execute(ctx *context.Context) (any, error)
}

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
