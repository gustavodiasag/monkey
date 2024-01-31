package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.IDENT, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "foo"},
					Value: "foo",
				},
			},
		},
	}

	if program.String() != "let var = foo;" {
		t.Errorf("program.String() wrong. got: %q", program.String())
	}
}
