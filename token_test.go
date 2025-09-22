package templatex_test

import (
	"github.com/ProImpact/templatex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenTypeToString(t *testing.T) {
	cases := []struct {
		tokenType templatex.TokenType
		expected  string
	}{
		{templatex.TEXT, "TEXT"},
		{templatex.LBRACE, "LBRACE"},
		{templatex.RBRACE, "RBRACE"},
		{templatex.LPARENT, "LPARENT"},
		{templatex.RPARENT, "RPARENT"},
		{templatex.LCORCH, "LCORCH"},
		{templatex.RCORCH, "RCORCH"},
		{templatex.INDENTIFIER, "INDENTIFIER"},
		{templatex.NUMBER, "NUMBER"},
		{templatex.DOBLE_EQUALS, "DOBLE_EQUALS"},
		{templatex.NOT_EQUALS, "NOT_EQUALS"},
		{templatex.GRATER, "GRATER"},
		{templatex.GRATER_EQUALS, "GRATER_EQUALS"},
		{templatex.LOWER_EQUALS, "LOWER_EQUALS"},
		{templatex.LOWER, "LOWER"},
		{templatex.COMMA, "COMMA"},
		{templatex.DOT, "DOT"},
		{templatex.EQUALS, "EQUALS"},
		{templatex.KEYWORD, "KEYWORD"},
		{templatex.ERROR, "ERROR"},
		{templatex.EOF, "EOF"},
	}

	for _, c := range cases {
		result := c.tokenType.String()
		assert.Equal(t, c.expected, result, "Expected %s, got %s", c.expected, result)
	}
}
