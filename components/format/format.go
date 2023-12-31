package format

import (
	"strings"

	"github.com/botscubes/bot-components/context"
)

func Format(str string, data *context.Context) (string, error) {
	runes := []rune(str)
	i := 0
	result := ""
	var prevRune rune = ' '
	for i < len(runes) {
		if prevRune == '\\' || (runes[i] != '{' && runes[i] != '\\' && runes[i] != '}') {
			result = result + string(runes[i])
		} else if runes[i] == '{' {
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
		} else if runes[i] == '}' {
			return "", ErrNoOpeningCurlyBrace
		}

		prevRune = runes[i]
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
