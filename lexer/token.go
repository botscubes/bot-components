package lexer

type TokenType = string

type Token struct {
	Type  TokenType
	Value string
}

type Position struct {
	Left  int
	Right int
}

func NewToken(tt TokenType, value string) *Token {
	return &Token{
		Type:  tt,
		Value: value,
	}
}

func NewPosition(left int, right int) *Position {
	return &Position{
		Left:  left,
		Right: right,
	}
}

const (
	ILLEGAL         = "ILLEGAL"
	INCOMPLETESTR   = "INCOMPLETESTR"   // "aaaa
	INCOMPLETEFLOAT = "INCOMPLETEFLOAT" // 12.
	NEGZERO         = "NEGZERO"         // -0
	EOF             = "EOF"

	IDENT  = "IDENT"  // x, t, struct
	INT    = "INT"    // 123 | -123
	FLOAT  = "FLOAT"  // 0.123 | -123.123
	STRING = "STRING" // "abcde"
	BOOL   = "BOOL"

	MINUS         = "-"
	EXCLAMINATION = "!"

	EQ  = "=="
	NEQ = "!="
	LEQ = "<="
	GEQ = ">="
	LT  = "<"
	GT  = ">"

	LAND = "&&"
	LOR  = "||"

	DOT = "."

	LPAR     = "("
	RPAR     = ")"
	LBRACKET = "["
	RBRACKET = "]"

	KEYWORD = "KEYWORD"
)
