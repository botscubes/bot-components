package format

import (
	"strings"

	"github.com/botscubes/bot-components/context"
)

func CheckFormatString(str string) error {
	runes := []rune(str)
	i := 0
	for i < len(runes) {
		if runes[i] == '\\' {
			i++
			if i >= len(runes) {
				return ErrUnknowEscapeSequence
			}
			if runes[i] != 'n' &&
				runes[i] != 't' &&
				runes[i] != '$' &&
				runes[i] != '\\' {
				return ErrUnknowEscapeSequence
			}
		} else if runes[i] == '$' {
			i++
			if i >= len(runes) {
				return ErrNoOpeningCurlyBrace
			}
			if runes[i] == '{' {

				i++
				r, err := getClosingCurlyBracePosition(runes, i)
				if err != nil {
					return err
				}
				err = context.CheckPath(strings.TrimSpace(string((runes[i:r]))))
				if err != nil {
					return err
				}
				i = r
			} else {
				return ErrNoOpeningCurlyBrace
			}
		}
		i++
	}

	return nil
}
