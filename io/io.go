package io

type ButtonData struct {
	Text string `json:"text"`
}

type IO interface {
	PrintText(text string) error
	PrintButtons(text string, buttons []*ButtonData) error
	PrintPhoto(file []byte, name string) error
	ReadText() *string
}
