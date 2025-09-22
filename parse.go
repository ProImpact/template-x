package templatex

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

var ErrInvalidToken = errors.New("invalid token")

func (t *templateParse) Parse() ([]*Node, error) {
	var nodes []*Node
	lineNumber := 0
	indexInline := 0

lopp:
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				t.appendEOF(&nodes)
				break lopp
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsSpace(char) {
			t.handleWhitespace(char, &lineNumber, &indexInline)
			continue
		}
		if unicode.IsDigit(char) {
			if err := t.readNumberToken(char, &nodes); err != nil {
				return nil, err
			}
			continue
		}
		if unicode.IsLetter(char) {
			if err := t.readIdentifierToken(char, &nodes); err != nil {
				return nil, err
			}
			continue
		}
		if err := t.handleSpecialCharacters(char, &nodes); err != nil {
			return nil, err
		}
	}

	return nodes, nil
}

func (t *templateParse) appendEOF(nodes *[]*Node) {
	*nodes = append(*nodes, &Node{
		NodeType: EOF,
		Lexema:   string("0"),
	})
}

func (t *templateParse) handleWhitespace(char rune, lineNumber, indexInline *int) {
	if char == '\n' {
		*lineNumber++
		*indexInline = 0
	} else {
		*indexInline++
	}
}

func (t *templateParse) readNumberToken(char rune, nodes *[]*Node) error {
	number := 0
	nm, _ := strconv.Atoi(string(char))
	number += nm
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: NUMBER,
					Lexema:   fmt.Sprintf("%d", number),
				})
				t.appendEOF(nodes)
				return nil
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsDigit(char) {
			nm, _ := strconv.Atoi(string(char))
			number = number*10 + nm
		} else {
			_ = t.buffer.UnreadRune()
			*nodes = append(*nodes, &Node{
				NodeType: NUMBER,
				Lexema:   fmt.Sprintf("%d", number),
			})
			break
		}
	}
	return nil
}

func (t *templateParse) readIdentifierToken(char rune, nodes *[]*Node) error {
	textIdentifier := bytes.NewBuffer([]byte(string(char)))
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: INDENTIFIER,
					Lexema:   textIdentifier.String(),
				})
				t.appendEOF(nodes)
				return nil
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsLetter(char) || char == '_' {
			_, err = textIdentifier.WriteRune(char)
			if err != nil {
				return fmt.Errorf("error writing rune: %w", err)
			}
		} else {
			_ = t.buffer.UnreadRune()
			*nodes = append(*nodes, &Node{
				NodeType: INDENTIFIER,
				Lexema:   textIdentifier.String(),
			})
			break
		}
	}
	return nil
}

func (t *templateParse) handleSpecialCharacters(char rune, nodes *[]*Node) error {
	switch char {
	case '.':
		*nodes = append(*nodes, &Node{NodeType: DOT, Lexema: string(char)})
	case '{':
		*nodes = append(*nodes, &Node{NodeType: LBRACE, Lexema: string(char)})
	// Add cases for other special characters
	case ')':
		*nodes = append(*nodes, &Node{
			NodeType: RPARENT,
			Lexema:   string(char),
		})
	case '[':
		*nodes = append(*nodes, &Node{
			NodeType: LCORCH,
			Lexema:   string(char),
		})
	case ']':
		*nodes = append(*nodes, &Node{
			NodeType: RCORCH,
			Lexema:   string(char),
		})
	case ',':
		*nodes = append(*nodes, &Node{
			NodeType: COMMA,
			Lexema:   string(char),
		})
	case '"':
		textBuffer := bytes.NewBuffer([]byte(""))
		for {
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return fmt.Errorf("recheated to the end of a file, expeting close text token '\"")
				}
				return fmt.Errorf("error reading the rune: %w", err)
			}
			_, err = textBuffer.WriteRune(char)
			if err != nil {
				return fmt.Errorf("err writing to the text buffer: %w", err)
			}
			if char == '"' {
				break
			}
		}
		*nodes = append(*nodes, &Node{
			NodeType: TEXT,
			Lexema:   textBuffer.String(),
		})
	case '=':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: EQUALS,
					Lexema:   string(char),
				})
				t.appendEOF(nodes)
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			*nodes = append(*nodes, &Node{
				NodeType: DOBLE_EQUALS,
				Lexema:   string(char),
			})
			return nil
		}
		*nodes = append(*nodes, &Node{
			NodeType: EQUALS,
			Lexema:   string(char),
		})
	case '<':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: LOWER,
					Lexema:   string(char),
				})
				*nodes = append(*nodes, &Node{
					NodeType: EOF,
					Lexema:   string("0"),
				})
				return nil
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			*nodes = append(*nodes, &Node{
				NodeType: LOWER_EQUALS,
				Lexema:   string(char),
			})
			return nil
		}
		*nodes = append(*nodes, &Node{
			NodeType: LOWER,
			Lexema:   string(char),
		})
	case '>':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: GRATER,
					Lexema:   string(char),
				})
				t.appendEOF(nodes)
				return nil
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			*nodes = append(*nodes, &Node{
				NodeType: GRATER_EQUALS,
				Lexema:   string(char),
			})
		}
		*nodes = append(*nodes, &Node{
			NodeType: GRATER,
			Lexema:   string(char),
		})
	case '!':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				*nodes = append(*nodes, &Node{
					NodeType: ERROR,
					Lexema:   "expected character = after !",
				})
				*nodes = append(*nodes, &Node{
					NodeType: EOF,
					Lexema:   string("0"),
				})
			}
			return fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			*nodes = append(*nodes, &Node{
				NodeType: NOT_EQUALS,
				Lexema:   string(char),
			})
			return nil
		}
		*nodes = append(*nodes, &Node{
			NodeType: ERROR,
			Lexema:   "expected character =",
		})
	default:
		return fmt.Errorf("unrecognized character: %c", char)
	}
	return nil
}

