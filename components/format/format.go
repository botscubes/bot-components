package format

import (
	"strings"

	"github.com/botscubes/bot-components/context"
)

func Format(str string, data *context.Context) (string, error) {
	runes := []rune(str)
	i := 0
	result := ""
	for i < len(runes) {
		if runes[i] == '\\' {
			i++
			if i >= len(runes) {
				return "", ErrUnknowEscapeSequence
			}
			if runes[i] == 'n' {
				result = result + string('\n')
			} else if runes[i] == 't' {
				result = result + string('\t')
			} else if runes[i] == '$' {
				result = result + string('$')
			} else if runes[i] == '\\' {
				result = result + string('\\')
			} else {
				return "", ErrUnknowEscapeSequence
			}
		} else if runes[i] == '$' {
			i++
			if i >= len(runes) {
				return "", ErrNoOpeningCurlyBrace
			}
			if runes[i] == '{' {

				i++
				r, err := getClosingCurlyBracePosition(runes, i)
				if err != nil {
					return "", err
				}
				v, err := data.GetValue(strings.TrimSpace(string((runes[i:r]))))
				if err != nil {
					return "", err
				}
				s, err := v.ToString()
				if err != nil {
					return "", err
				}
				result = result + s
				i = r
			} else {
				return "", ErrNoOpeningCurlyBrace
			}
		} else if runes[i] != '$' {
			result = result + string(runes[i])
		}
		i++
	}

	return result, nil
}

func getClosingCurlyBracePosition(runes []rune, l int) (int, error) {
	r := l
	for {
		if r < len(runes) {
			if runes[r] == '}' {
				return r, nil
			}
		} else {
			return 0, ErrNoClosingCurlyBrace
		}
		r++
	}

}
