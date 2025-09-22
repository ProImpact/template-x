package templatex

import "slices"

type TokenType int

const (
	_             TokenType = iota
	TEXT                    // "name"
	LBRACE                  // {
	RBRACE                  // }
	LPARENT                 // (
	RPARENT                 // )
	LCORCH                  // [
	RCORCH                  // ]
	INDENTIFIER             // variable name
	NUMBER                  // 21, 40
	DOBLE_EQUALS            // ==
	NOT_EQUALS              // !=
	GRATER                  // >
	GRATER_EQUALS           // >=
	LOWER_EQUALS            // <=
	LOWER                   // <
	COMMA                   // ,
	DOT                     // .
	EQUALS                  // =
	KEYWORD
	ERROR
	EOF
)

type Keyword = string

const (
	IF = iota
	ELSE_IF
	ELSE
	FOR
	IN
	OR
	AND
	VAR
)

var Keywords = []Keyword{
	IF:      "if",
	ELSE_IF: "else if",
	ELSE:    "else",
	FOR:     "for",
	IN:      "in",
	OR:      "or",
	AND:     "and",
	VAR:     "var",
}

type Node struct {
	NodeType TokenType
	Lexema   string
}

func IsKeyword(keyword string) bool {
	return slices.Contains(Keywords, keyword)
}
