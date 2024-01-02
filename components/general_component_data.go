package components

type GeneralComponentData struct {
	Type            ComponentType `json:"componentType"`
	NextComponentId int           `json:"nextComponentId"`
	SavePath        string        `json:"savePath"`
}
