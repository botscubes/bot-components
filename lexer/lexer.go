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

func (lx *Lexer) NextToken() *Token {
	var token *Token

	//	ILLEGAL = "ILLEGAL"
	//
	//	IDENT  = "IDENT"  // x, t, struct
	//	INT    = "INT"    // 123
	//	STRING = "STRING" // "abcde"

	//	EQ  = "=="
	//	NEQ = "!="
	//	LEQ = "<="
	//	GEQ = ">="
	//	LT  = "<"
	//	GT  = ">"
	//
	//	LAND = "&&"
	//	LOR  = "||"
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
		token = NewToken(MINUS, MINUS)
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
		if isLetter(ch) {
			token = lx.getIdentToken(pos)
		} else if isDigit(ch) {
			token = lx.getUintToken(pos)
		} else {
			token = NewToken(ILLEGAL, string(ch))
		}
	}

	return token
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

func (lx *Lexer) getUintToken(l int) *Token {
	if lx.str[l] == '0' {
		return NewToken(UINT, "0")
	}
	for {
		ch, pos := lx.char()
		if !isDigit(ch) {
			return NewToken(UINT, lx.str[l:pos])
		}
		lx.pos++
	}
}

func (lx *Lexer) getIdentToken(l int) *Token {
	for {
		ch, pos := lx.char()
		if !isDigit(ch) && !isLetter(ch) && ch != '_' {
			return NewToken(IDENT, lx.str[l:pos])
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
