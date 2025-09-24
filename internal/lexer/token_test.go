package lexer_test

import (
	"testing"

	"github.com/ProImpact/templatex/internal/lexer"
	"github.com/stretchr/testify/assert"
)

func TestTokenTypeToString(t *testing.T) {
	cases := []struct {
		tokenType lexer.TokenType
		expected  string
	}{
		{lexer.TEXT, "TEXT"},
		{lexer.LBRACE, "LBRACE"},
		{lexer.RBRACE, "RBRACE"},
		{lexer.LPARENT, "LPARENT"},
		{lexer.RPARENT, "RPARENT"},
		{lexer.LCORCH, "LCORCH"},
		{lexer.RCORCH, "RCORCH"},
		{lexer.INDENTIFIER, "INDENTIFIER"},
		{lexer.NUMBER, "NUMBER"},
		{lexer.DOBLE_EQUALS, "DOBLE_EQUALS"},
		{lexer.NOT_EQUALS, "NOT_EQUALS"},
		{lexer.GRATER, "GRATER"},
		{lexer.GRATER_EQUALS, "GRATER_EQUALS"},
		{lexer.LOWER_EQUALS, "LOWER_EQUALS"},
		{lexer.LOWER, "LOWER"},
		{lexer.COMMA, "COMMA"},
		{lexer.DOT, "DOT"},
		{lexer.EQUALS, "EQUALS"},
		{lexer.KEYWORD, "KEYWORD"},
		{lexer.ERROR, "ERROR"},
		{lexer.EOF, "EOF"},
	}

	for _, c := range cases {
		result := c.tokenType.String()
		assert.Equal(t, c.expected, result, "Expected %s, got %s", c.expected, result)
	}
}
