package format

import (
	"github.com/botscubes/bot-components/components"
)

type FormatComponent struct {
	FormatString string `json:"formatString"`
	components.GeneralComponentData
}
