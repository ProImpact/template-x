package ast

import (
	"github.com/ProImpact/templatex/internal/lexer"
)

type AST struct {
	lex *lexer.Lexer
}

func MustNewAST(filePath string) *AST {
	lex, err := lexer.NewLexer(filePath, true)
	if err != nil {
		panic(err)
	}
	return &AST{
		lex: lex,
	}
}
