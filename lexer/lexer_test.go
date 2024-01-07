package lexer

import "testing"

type comparison struct {
	str    string
	tokens []*Token
}

var comparsions []*comparison = []*comparison{
	{
		str: "a && b || ccc>=d || !(n!=123.23 && m <=  -0.1265)",
		tokens: []*Token{
			NewToken(IDENT, "a"),
			NewToken(LAND, LAND),
			NewToken(IDENT, "b"),
			NewToken(LOR, LOR),
			NewToken(IDENT, "ccc"),
			NewToken(GEQ, GEQ),
			NewToken(IDENT, "d"),
			NewToken(LOR, LOR),
			NewToken(EXCLAMINATION, EXCLAMINATION),
			NewToken(LPAR, LPAR),
			NewToken(IDENT, "n"),
			NewToken(NEQ, NEQ),
			NewToken(FLOAT, "123.23"),
			NewToken(LAND, LAND),
			NewToken(IDENT, "m"),
			NewToken(LEQ, LEQ),
			NewToken(FLOAT, "-0.1265"),
			NewToken(RPAR, RPAR),
		},
	},
	{
		str: "\"ффаааппппв\" \"ффффываы ",
		tokens: []*Token{
			NewToken(STRING, "ффаааппппв"),
			NewToken(INCOMPLETESTR, "ффффываы "),
		},
	},
}

func TestLexer(t *testing.T) {
	for _, c := range comparsions {
		lx := NewLexer(c.str)
		tk, _ := lx.NextToken()
		i := 0
		for tk.Type != EOF {
			t.Logf("Token: type: %s, value: %s", tk.Type, tk.Value)
			if i >= len(c.tokens) {
				t.Fatalf("Number of elements does not match, token: type: %s, value: %s", tk.Type, tk.Value)
			}
			ctk := c.tokens[i]
			if tk.Type != ctk.Type {
				t.Fatalf("Types don't match: %s : %s", tk.Type, ctk.Type)
			}
			if tk.Value != ctk.Value {
				t.Fatalf("Values don't match: %s : %s", tk.Value, ctk.Value)
			}
			tk, _ = lx.NextToken()
			i++
		}
	}
}
