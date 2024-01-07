package lexer

import "testing"

type comparison struct {
	str    string
	tokens []*Token
}

var comparsions []*comparison = []*comparison{
	{
		str: "a && b || ccc>=d",
		tokens: []*Token{
			NewToken(IDENT, "a"),
			NewToken(LAND, LAND),
			NewToken(IDENT, "b"),
			NewToken(LOR, LOR),
			NewToken(IDENT, "ccc"),
			NewToken(GEQ, GEQ),
			NewToken(IDENT, "d"),
		},
	},
}

func TestLexer(t *testing.T) {
	for _, c := range comparsions {
		lx := NewLexer(c.str)
		tk := lx.NextToken()
		i := 0
		for tk.Type != EOF {
			if i >= len(c.tokens) {
				t.Fatal("Number of elements does not match")
			}
			ctk := c.tokens[i]
			if tk.Type != ctk.Type {
				t.Fatalf("Types don't match: %s : %s", tk.Type, ctk.Type)
			}
			if tk.Value != ctk.Value {
				t.Fatalf("Values don't match: %s : %s", tk.Value, ctk.Value)
			}
			tk = lx.NextToken()
			i++
		}
	}
}
