package io

type IO interface {
	OutputText(text string)
	InputText() *string
}
