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

func (t *templateParse) Next() (*Node, error) {
	lineNumber := 0
	indexInline := 0

	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return &Node{
					NodeType: EOF,
					Lexema:   "0",
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsSpace(char) {
			t.handleWhitespace(char, &lineNumber, &indexInline)
			continue
		}
		if unicode.IsDigit(char) {
			num, err := t.readNumberToken(char)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					return nil, err
				}
			}
			return &Node{
				NodeType: NUMBER,
				Lexema:   fmt.Sprintf("%d", num),
			}, nil
		}
		if unicode.IsLetter(char) {
			token, err := t.readIdentifierToken(char)
			if err != nil {
				return nil, err
			}
			return token, nil
		}
		token, err := t.handleSpecialCharacters(char)
		if err != nil {
			return nil, err
		}
		return token, nil
	}
}

func (t *templateParse) handleWhitespace(char rune, lineNumber, indexInline *int) {
	if char == '\n' {
		*lineNumber++
		*indexInline = 0
	} else {
		*indexInline++
	}
}

func (t *templateParse) readNumberToken(initialChar rune) (int, error) {
	number := 0
	nm, _ := strconv.Atoi(string(initialChar))
	number += nm
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return number, t.buffer.UnreadRune()
			}
			return -1, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsDigit(char) {
			nm, _ := strconv.Atoi(string(char))
			number = number*10 + nm
		} else {
			err = t.buffer.UnreadRune()
			if err != nil {
				return -1, err
			}
			break
		}
	}
	return number, nil
}

func (t *templateParse) readIdentifierToken(initialChar rune) (*Node, error) {
	textIdentifier := bytes.NewBuffer([]byte(string(initialChar)))
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if IsKeyword(textIdentifier.String()) {
					return &Node{
						NodeType: KEYWORD,
						Lexema:   textIdentifier.String(),
					}, nil

				} else {
					return &Node{
						NodeType: INDENTIFIER,
						Lexema:   textIdentifier.String(),
					}, nil
				}
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsLetter(char) || char == '_' {
			_, err = textIdentifier.WriteRune(char)
			if err != nil {
				return nil, fmt.Errorf("error writing rune: %w", err)
			}
		} else {
			_ = t.buffer.UnreadRune()
			if IsKeyword(textIdentifier.String()) {
				return &Node{
					NodeType: KEYWORD,
					Lexema:   textIdentifier.String(),
				}, nil

			} else {
				return &Node{
					NodeType: INDENTIFIER,
					Lexema:   textIdentifier.String(),
				}, nil
			}
		}
	}
}

func (t *templateParse) handleSpecialCharacters(char rune) (*Node, error) {
	switch char {
	case '.':
		return &Node{NodeType: DOT, Lexema: string(char)}, nil
	case '{':
		return &Node{NodeType: LBRACE, Lexema: string(char)}, nil
	// Add cases for other special characters
	case ')':
		return &Node{
			NodeType: RPARENT,
			Lexema:   string(char),
		}, nil
	case '(':
		return &Node{
			NodeType: LPARENT,
			Lexema:   string(char),
		}, nil
	case '|':
		return &Node{
			NodeType: LINE,
			Lexema:   string(char),
		}, nil
	case '/':
		return &Node{
			NodeType: SLASH,
			Lexema:   string(char),
		}, nil
	case '*':
		return &Node{
			NodeType: ASTERISIC,
			Lexema:   string(char),
		}, nil
	case '$':
		return &Node{
			NodeType: DOLLAR,
			Lexema:   string(char),
		}, nil
	case ':':
		return &Node{
			NodeType: COLON,
			Lexema:   string(char),
		}, nil
	case '}':
		return &Node{
			NodeType: RBRACE,
			Lexema:   string(char),
		}, nil
	case '[':
		return &Node{
			NodeType: LCORCH,
			Lexema:   string(char),
		}, nil
	case ']':
		return &Node{
			NodeType: RCORCH,
			Lexema:   string(char),
		}, nil
	case ',':
		return &Node{
			NodeType: COMMA,
			Lexema:   string(char),
		}, nil
	case '"':
		textBuffer := bytes.NewBuffer([]byte("\""))
		for {
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil, fmt.Errorf("recheated to the end of a file, expeting close text token '\"")
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			_, err = textBuffer.WriteRune(char)
			if err != nil {
				return nil, fmt.Errorf("err writing to the text buffer: %w", err)
			}
			if char == '"' {
				break
			}
		}
		return &Node{
			NodeType: TEXT,
			Lexema:   textBuffer.String(),
		}, nil
	case '=':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = t.buffer.UnreadRune()
				if err != nil {
					return nil, err
				}
				return &Node{
					NodeType: EQUALS,
					Lexema:   string(char),
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			return &Node{
				NodeType: DOBLE_EQUALS,
				Lexema:   string("=="),
			}, nil
		}
		return &Node{
			NodeType: EQUALS,
			Lexema:   string(char),
		}, nil
	case '<':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = t.buffer.UnreadRune()
				if err != nil {
					return nil, err
				}
				return &Node{
					NodeType: LOWER,
					Lexema:   string(char),
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			return &Node{
				NodeType: LOWER_EQUALS,
				Lexema:   string(char),
			}, nil
		}
		return &Node{
			NodeType: LOWER,
			Lexema:   string(char),
		}, nil
	case '>':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = t.buffer.UnreadRune()
				if err != nil {
					return nil, err
				}
				return &Node{
					NodeType: GRATER,
					Lexema:   string(char),
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			return &Node{
				NodeType: GRATER_EQUALS,
				Lexema:   string(char),
			}, nil
		}
		return &Node{
			NodeType: GRATER,
			Lexema:   string(char),
		}, nil
	case '!':
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return &Node{
					NodeType: ERROR,
					Lexema:   "expected character = after !",
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			return &Node{
				NodeType: NOT_EQUALS,
				Lexema:   string(char),
			}, nil
		}
		return &Node{
			NodeType: ERROR,
			Lexema:   "expected character =",
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized character: %c", char)
	}
}
