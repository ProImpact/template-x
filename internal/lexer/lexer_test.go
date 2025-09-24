package lexer_test

import (
	"testing"

	"github.com/ProImpact/templatex/internal/lexer"
)

func TestGetTokenMetadata(t *testing.T) {
	lex, err := lexer.NewLexer("../../test/conditional.tmpl", false)
	if err != nil {
		t.Fatal(err)
	}
	for {
		lex.Advance()
		tok := lex.Current()
		if tok.NodeType == lexer.EOF {
			break
		}
		t.Logf("%+v \n", tok)
	}
}
