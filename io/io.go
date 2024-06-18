package io

type ButtonData struct {
	Text string `json:"text"`
}

type IO interface {
	PrintText(text string)
	PrintButtons(text string, buttons []*ButtonData)
	PrintPhoto(file []byte, name string) error
	ReadText() *string
}
