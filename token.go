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

func (t TokenType) String() string {
	switch t {
	case TEXT:
		return "TEXT"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LPARENT:
		return "LPARENT"
	case RPARENT:
		return "RPARENT"
	case LCORCH:
		return "LCORCH"
	case RCORCH:
		return "RCORCH"
	case INDENTIFIER:
		return "INDENTIFIER"
	case NUMBER:
		return "NUMBER"
	case DOBLE_EQUALS:
		return "DOBLE_EQUALS"
	case NOT_EQUALS:
		return "NOT_EQUALS"
	case GRATER:
		return "GRATER"
	case GRATER_EQUALS:
		return "GRATER_EQUALS"
	case LOWER_EQUALS:
		return "LOWER_EQUALS"
	case LOWER:
		return "LOWER"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case EQUALS:
		return "EQUALS"
	case KEYWORD:
		return "KEYWORD"
	case ERROR:
		return "ERROR"
	case EOF:
		return "EOF"
	default:
		return "UNKNOWN"
	}
}

func IsKeyword(keyword string) bool {
	return slices.Contains(Keywords, keyword)
}
