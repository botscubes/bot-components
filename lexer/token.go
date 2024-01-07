package lexer

type TokenType = string

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tt TokenType, value string) *Token {
	return &Token{
		Type:  tt,
		Value: value,
	}
}

const (
	ILLEGAL       = "ILLEGAL"
	INCOMPLETESTR = "INCOMPLETESTR" // "aaaa
	EOF           = "EOF"

	IDENT  = "IDENT"  // x, t, struct
	UINT   = "UINT"   // 123
	STRING = "STRING" // "abcde"

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

	TRUE  = "TRUE"
	FALSE = "FALSE"
)
