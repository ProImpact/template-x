package lexer

import "slices"

type TokenType int

const (
	_             TokenType = iota
	TEXT                    // "name"
	SLASH                   // /
	ASTERISIC               // *
	DOLLAR                  // $
	COLON                   // :
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
	LINE
	ERROR
	EOF
)

type Keyword = string

const (
	IF = iota
	ELSE
	FOR
	IN
	OR
	AND
	VAR
	DEFINE
	TEMPLATE
	END
	FALSE
	TRUE
	WITH
	BLOCK
)

var Keywords = []Keyword{
	IF:       "if",
	ELSE:     "else",
	FOR:      "for",
	IN:       "in",
	OR:       "or",
	AND:      "and",
	VAR:      "var",
	DEFINE:   "define",
	TEMPLATE: "template",
	END:      "end",
	FALSE:    "false",
	TRUE:     "true",
	WITH:     "with",
	BLOCK:    "block",
}

type Node struct {
	NodeType   TokenType
	Lexema     string
	LineIndex  int
	LineNumber int
}

func (t TokenType) String() string {
	switch t {
	case LINE:
		return "LINE"
	case SLASH:
		return "SLASH"
	case ASTERISIC:
		return "ASTERISIC"
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
	case COLON:
		return "COLON"
	case DOLLAR:
		return "DOLLAR"
	default:
		return "UNKNOWN"
	}
}

func IsKeyword(keyword string) bool {
	return slices.Contains(Keywords, keyword)
}
