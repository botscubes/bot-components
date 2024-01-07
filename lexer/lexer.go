package lexer

type Lexer struct {
	str string
	pos int
}

func NewLexer(str string) *Lexer {
	return &Lexer{
		str: str,
		pos: 0,
	}
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func (lx *Lexer) NextToken() (*Token, *Position) {
	var token *Token

	lx.skipSpaces()
	ch, pos := lx.nextChar()
	switch ch {
	case 0:
		token = NewToken(EOF, EOF)
	case '(':
		token = NewToken(LPAR, LPAR)
	case ')':
		token = NewToken(RPAR, RPAR)
	case '[':
		token = NewToken(LBRACKET, LBRACKET)
	case ']':
		token = NewToken(RBRACKET, RBRACKET)
	case '.':
		token = NewToken(DOT, DOT)
	case '!':
		lc, _ := lx.char()
		if lc == '=' {
			lx.pos++
			token = NewToken(NEQ, NEQ)
		} else {
			token = NewToken(EXCLAMINATION, EXCLAMINATION)
		}
	case '-':
		lc, _ := lx.char()
		if isDigit(lc) {
			token = lx.getNumberToken(pos)
		} else {
			token = NewToken(MINUS, MINUS)
		}
	case '|':
		lc, _ := lx.char()
		if lc == '|' {
			lx.pos++
			token = NewToken(LOR, LOR)
		} else {

			token = NewToken(ILLEGAL, string(ch))
		}
	case '&':
		lc, _ := lx.char()
		if lc == '&' {
			lx.pos++
			token = NewToken(LAND, LAND)
		} else {
			token = NewToken(ILLEGAL, string(ch))
		}
	case '"':
		token = lx.getStringToken()
	case '=':
		lc, _ := lx.char()
		if lc == '=' {
			lx.pos++
			token = NewToken(EQ, EQ)
		} else {
			token = NewToken(ILLEGAL, string(ch))
		}
	case '>':
		lc, _ := lx.char()
		if lc == '=' {
			lx.pos++
			token = NewToken(GEQ, GEQ)
		} else {
			token = NewToken(GT, GT)
		}
	case '<':
		lc, _ := lx.char()
		if lc == '=' {
			lx.pos++
			token = NewToken(LEQ, LEQ)
		} else {
			token = NewToken(LT, LT)
		}
	default:
		if isLetter(ch) || ch == '_' {
			token = lx.getIdentOrKeywordToken(pos)
		} else if isDigit(ch) {
			token = lx.getNumberToken(pos)
		} else {
			token = NewToken(ILLEGAL, string(ch))
		}
	}

	return token, NewPosition(pos, lx.pos)
}
func (lx *Lexer) nextChar() (byte, int) {
	if lx.isNextChar() {
		ch, pos := lx.char()
		lx.pos++
		return ch, pos
	}
	return 0, 0
}
func (lx *Lexer) char() (byte, int) {
	if lx.isNextChar() {
		return lx.str[lx.pos], lx.pos
	}
	return 0, lx.pos
}

func (lx *Lexer) isNextChar() bool {
	return lx.pos < len(lx.str)
}

func (lx *Lexer) getStringToken() *Token {
	l := lx.pos
	for {
		ch, pos := lx.nextChar()
		if ch == 0 {
			return NewToken(INCOMPLETESTR, lx.str[l:pos])
		} else if ch == '"' {
			return NewToken(STRING, lx.str[l:pos])
		}
	}
}

func (lx *Lexer) getNumberToken(l int) *Token {
	if lx.str[l] == '-' {
		ch, _ := lx.char()
		if ch == '0' {
			lx.pos++
			ch, _ := lx.char()
			if ch == '.' {
				lx.pos++
				ch, pos := lx.char()
				if !isDigit(ch) {
					return NewToken(INCOMPLETEFLOAT, lx.str[l:pos])
				}

				return NewToken(FLOAT, lx.str[l:pos]+lx.getFraction(pos))
			}
			return NewToken(NEGZERO, NEGZERO)
		}
	}
	if lx.str[l] == '0' {
		ch, _ := lx.char()
		if ch == '.' {
			lx.pos++
			ch, pos := lx.char()
			if !isDigit(ch) {
				return NewToken(INCOMPLETEFLOAT, lx.str[l:pos])
			}

			return NewToken(FLOAT, lx.str[l:pos]+lx.getFraction(pos))
		}

		return NewToken(INT, "0")

	}

	for {
		ch, pos := lx.char()
		if ch == '.' {
			lx.pos++
			ch, pos := lx.char()
			if !isDigit(ch) {
				return NewToken(INCOMPLETEFLOAT, lx.str[l:pos])
			}
			return NewToken(FLOAT, lx.str[l:pos]+lx.getFraction(pos))
		} else if !isDigit(ch) {
			return NewToken(INT, lx.str[l:pos])
		}
		lx.pos++
	}
}

func (lx *Lexer) getFraction(l int) string {
	for {
		ch, pos := lx.char()
		if !isDigit(ch) {
			return lx.str[l:pos]
		}
		lx.pos++
	}
}

func (lx *Lexer) getIdentOrKeywordToken(l int) *Token {
	for {
		ch, pos := lx.char()
		if !isDigit(ch) && !isLetter(ch) && ch != '_' {
			str := lx.str[l:pos]
			switch str {
			case "true":
				return NewToken(BOOL, str)
			case "false":
				return NewToken(BOOL, str)
			}
			return NewToken(IDENT, str)
		}
		lx.pos++
	}
}

func (lx *Lexer) skipSpaces() {
	for {
		ch, _ := lx.char()

		if ch != ' ' {
			return
		}
		lx.pos++
	}
}
