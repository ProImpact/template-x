package lexer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"unicode"
)

var ErrInvalidToken = errors.New("invalid token")

func (t *Lexer) nextToken() (*Node, error) {
	token, err := t.next()
	if err != nil {
		return nil, err
	}
	t.previous = token
	return token, nil
}

func (t *Lexer) Previous() *Node {
	return t.previous
}

func (t *Lexer) next() (*Node, error) {

	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				t.LineIndex++
				return &Node{
					NodeType:   EOF,
					Lexema:     "0",
					LineIndex:  t.LineIndex,
					LineNumber: t.LineNumber,
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsSpace(char) {
			t.handleWhitespace(char, &t.LineNumber, &t.LineIndex)
			continue
		}
		if unicode.IsDigit(char) {
			num, initialPos, err := t.readNumberToken(char)
			if err != nil {
				if !errors.Is(err, io.EOF) {
					return nil, err
				}
			}
			return &Node{
				NodeType:   NUMBER,
				Lexema:     fmt.Sprintf("%d", num),
				LineIndex:  initialPos,
				LineNumber: t.LineNumber,
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

func (t *Lexer) handleWhitespace(char rune, lineNumber, indexInline *int) {
	if char == '\n' {
		*lineNumber++
		*indexInline = 0
	} else {
		*indexInline++
	}
}

func (t *Lexer) readNumberToken(initialChar rune) (int, int, error) {
	t.LineIndex++
	initialPos := t.LineIndex
	number := 0
	nm, _ := strconv.Atoi(string(initialChar))
	number += nm
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return number, initialPos, t.buffer.UnreadRune()
			}
			return -1, initialPos, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsDigit(char) {
			nm, _ := strconv.Atoi(string(char))
			number = number*10 + nm
			t.LineIndex++
		} else {
			err = t.buffer.UnreadRune()
			if err != nil {
				return -1, initialPos, err
			}
			break
		}
	}
	return number, initialPos, nil
}

func (t *Lexer) readIdentifierToken(initialChar rune) (*Node, error) {
	textIdentifier := bytes.NewBuffer([]byte(string(initialChar)))
	t.LineIndex++
	initialPost := t.LineIndex
	for {
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if IsKeyword(textIdentifier.String()) {
					return &Node{
						NodeType:   KEYWORD,
						Lexema:     textIdentifier.String(),
						LineIndex:  initialPost,
						LineNumber: t.LineNumber,
					}, nil

				} else {
					return &Node{
						NodeType:   INDENTIFIER,
						Lexema:     textIdentifier.String(),
						LineIndex:  initialPost,
						LineNumber: t.LineNumber,
					}, nil
				}
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if unicode.IsLetter(char) || char == '_' {
			_, err = textIdentifier.WriteRune(char)
			t.LineIndex++
			if err != nil {
				return nil, fmt.Errorf("error writing rune: %w", err)
			}
		} else {
			_ = t.buffer.UnreadRune()
			if IsKeyword(textIdentifier.String()) {
				return &Node{
					NodeType:   KEYWORD,
					Lexema:     textIdentifier.String(),
					LineIndex:  initialPost,
					LineNumber: t.LineNumber,
				}, nil

			} else {
				return &Node{
					NodeType:   INDENTIFIER,
					Lexema:     textIdentifier.String(),
					LineIndex:  initialPost,
					LineNumber: t.LineNumber,
				}, nil
			}
		}
	}
}

func (t *Lexer) handleSpecialCharacters(char rune) (*Node, error) {
	switch char {
	case '.':
		t.LineIndex++
		return &Node{NodeType: DOT, Lexema: string(char), LineIndex: t.LineIndex, LineNumber: t.LineNumber}, nil
	case '@':
		t.LineIndex++
		return &Node{NodeType: ARROBA, Lexema: string(char), LineIndex: t.LineIndex, LineNumber: t.LineNumber}, nil
	case '{':
		t.LineIndex++
		return &Node{NodeType: LBRACE, Lexema: string(char), LineIndex: t.LineIndex, LineNumber: t.LineNumber}, nil
	// Add cases for other special characters
	case ')':
		t.LineIndex++
		return &Node{
			NodeType:   RPARENT,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '(':
		t.LineIndex++
		return &Node{
			NodeType:   LPARENT,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '|':
		t.LineIndex++
		return &Node{
			NodeType:   LINE,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '/':
		t.LineIndex++
		return &Node{
			NodeType:   SLASH,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '*':
		t.LineIndex++
		return &Node{
			NodeType:   ASTERISIC,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '$':
		t.LineIndex++
		return &Node{
			NodeType:   DOLLAR,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case ':':
		t.LineIndex++
		return &Node{
			NodeType:   COLON,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '}':
		t.LineIndex++
		return &Node{
			NodeType:   RBRACE,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '[':
		t.LineIndex++
		return &Node{
			NodeType:   LCORCH,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case ']':
		t.LineIndex++
		return &Node{
			NodeType:   RCORCH,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case ',':
		t.LineIndex++
		return &Node{
			NodeType:   COMMA,
			Lexema:     string(char),
			LineIndex:  t.LineIndex,
			LineNumber: t.LineNumber,
		}, nil
	case '"':
		t.LineIndex++
		initialPos := t.LineIndex
		textBuffer := bytes.NewBuffer([]byte("\""))
		for {
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					return nil, fmt.Errorf("recheated to the end of a file, expeting close text token '\"")
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			t.LineIndex++
			_, err = textBuffer.WriteRune(char)
			if err != nil {
				return nil, fmt.Errorf("err writing to the text buffer: %w", err)
			}
			if char == '"' {
				break
			}
		}
		return &Node{
			NodeType:   TEXT,
			Lexema:     textBuffer.String(),
			LineIndex:  initialPos,
			LineNumber: t.LineNumber,
		}, nil
	case '=':
		t.LineIndex++
		initialPos := t.LineIndex
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = t.buffer.UnreadRune()
				if err != nil {
					return nil, err
				}
				return &Node{
					NodeType:   EQUALS,
					Lexema:     string(char),
					LineIndex:  initialPos,
					LineNumber: t.LineNumber,
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			t.LineIndex++
			return &Node{
				NodeType:   DOBLE_EQUALS,
				Lexema:     string("=="),
				LineIndex:  initialPos,
				LineNumber: t.LineNumber,
			}, nil
		}
		return &Node{
			NodeType:   EQUALS,
			Lexema:     string(char),
			LineIndex:  initialPos,
			LineNumber: t.LineNumber,
		}, nil
	case '<':
		t.LineIndex++
		initialPos := t.LineIndex
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				err = t.buffer.UnreadRune()
				if err != nil {
					return nil, err
				}
				return &Node{
					NodeType:   LOWER,
					Lexema:     string(char),
					LineIndex:  initialPos,
					LineNumber: t.LineNumber,
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			t.LineIndex++
			return &Node{
				NodeType:   LOWER_EQUALS,
				Lexema:     string(char),
				LineIndex:  initialPos,
				LineNumber: t.LineNumber,
			}, nil
		}
		return &Node{
			NodeType:   LOWER,
			Lexema:     string(char),
			LineIndex:  initialPos,
			LineNumber: t.LineNumber,
		}, nil
	case '>':
		t.LineIndex++
		initialPos := t.LineIndex
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
			t.LineIndex++
			return &Node{
				NodeType:   GRATER_EQUALS,
				Lexema:     string(char),
				LineIndex:  initialPos,
				LineNumber: t.LineNumber,
			}, nil
		}
		return &Node{
			NodeType:   GRATER,
			Lexema:     string(char),
			LineIndex:  initialPos,
			LineNumber: t.LineNumber,
		}, nil
	case '!':
		t.LineIndex++
		initialPos := t.LineIndex
		char, _, err := t.buffer.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return &Node{
					NodeType:   ERROR,
					Lexema:     "expected character = after !",
					LineIndex:  initialPos,
					LineNumber: t.LineNumber,
				}, nil
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '=' {
			t.LineIndex++
			return &Node{
				NodeType:   NOT_EQUALS,
				Lexema:     string(char),
				LineIndex:  initialPos,
				LineNumber: t.LineNumber,
			}, nil
		}
		return &Node{
			NodeType:   ERROR,
			Lexema:     "expected character =",
			LineIndex:  initialPos,
			LineNumber: t.LineNumber,
		}, nil
	default:
		return nil, fmt.Errorf("unrecognized character: %c", char)
	}
}
