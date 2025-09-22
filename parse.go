package templatex

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
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
				nodes = append(nodes, &Node{
					NodeType: EOF,
					Lexema:   string("0"),
				})
				break lopp
			}
			return nil, fmt.Errorf("error reading the rune: %w", err)
		}
		if char == '\n' {
			lineNumber++
			indexInline = 0
			continue
		}
		if unicode.IsSpace(char) {
			indexInline++
			continue
		}
		if unicode.IsDigit(char) {
			number := 0
			nm, _ := strconv.Atoi(string(char))
			number += nm
			for {
				char, _, err = t.buffer.ReadRune()
				if err != nil {
					if errors.Is(err, io.EOF) {
						nodes = append(nodes, &Node{
							NodeType: NUMBER,
							Lexema:   fmt.Sprintf("%d", number),
						})
						nodes = append(nodes, &Node{
							NodeType: EOF,
							Lexema:   string("0"),
						})
						break lopp
					}
					return nil, fmt.Errorf("error reading the rune: %w", err)
				}
				if unicode.IsDigit(char) {
					nm, _ := strconv.Atoi(string(char))
					number *= 10
					number += nm
				} else {
					err = t.buffer.UnreadRune()
					if err != nil {
						log.Fatal(err)
					}
					nodes = append(nodes, &Node{
						NodeType: NUMBER,
						Lexema:   fmt.Sprintf("%d", number),
					})
					break
				}
			}
		}
		if unicode.IsLetter(char) {
			textIdentifier := bytes.NewBuffer([]byte(string(char)))
			for {
				char, _, err = t.buffer.ReadRune()
				if err != nil {
					if errors.Is(err, io.EOF) {
						nodes = append(nodes, &Node{
							NodeType: INDENTIFIER,
							Lexema:   textIdentifier.String(),
						})
						nodes = append(nodes, &Node{
							NodeType: EOF,
							Lexema:   string("0"),
						})
						break lopp
					}
					return nil, fmt.Errorf("error reading the rune: %w", err)
				}
				if unicode.IsLetter(char) || char == '_' {
					_, err = textIdentifier.WriteRune(char)
					if err != nil {
						log.Fatal(err)
					}
				} else {
					err = t.buffer.UnreadRune()
					if err != nil {
						log.Fatal(err)
					}
					nodes = append(nodes, &Node{
						NodeType: INDENTIFIER,
						Lexema:   textIdentifier.String(),
					})
					break
				}
			}
		}
		switch char {
		case '.':
			nodes = append(nodes, &Node{
				NodeType: DOT,
				Lexema:   string(char),
			})
		case '{':
			nodes = append(nodes, &Node{
				NodeType: LBRACE,
				Lexema:   string(char),
			})
		case '}':
			nodes = append(nodes, &Node{
				NodeType: RBRACE,
				Lexema:   string(char),
			})
		case '(':
			nodes = append(nodes, &Node{
				NodeType: LPARENT,
				Lexema:   string(char),
			})
		case ')':
			nodes = append(nodes, &Node{
				NodeType: RPARENT,
				Lexema:   string(char),
			})
		case '[':
			nodes = append(nodes, &Node{
				NodeType: LCORCH,
				Lexema:   string(char),
			})
		case ']':
			nodes = append(nodes, &Node{
				NodeType: RCORCH,
				Lexema:   string(char),
			})
		case ',':
			nodes = append(nodes, &Node{
				NodeType: COMMA,
				Lexema:   string(char),
			})
		case '"':
			textBuffer := bytes.NewBuffer([]byte(""))
			for {
				char, _, err = t.buffer.ReadRune()
				if err != nil {
					if errors.Is(err, io.EOF) {
						return nil, fmt.Errorf("recheated to the end of a file, expeting close text token '\"'")
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
			nodes = append(nodes, &Node{
				NodeType: TEXT,
				Lexema:   textBuffer.String(),
			})
		case '=':
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					nodes = append(nodes, &Node{
						NodeType: EQUALS,
						Lexema:   string(char),
					})
					nodes = append(nodes, &Node{
						NodeType: EOF,
						Lexema:   string("0"),
					})
					break lopp
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			if char == '=' {
				nodes = append(nodes, &Node{
					NodeType: DOBLE_EQUALS,
					Lexema:   string(char),
				})
				continue
			}
			nodes = append(nodes, &Node{
				NodeType: EQUALS,
				Lexema:   string(char),
			})
		case '<':
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					nodes = append(nodes, &Node{
						NodeType: LOWER,
						Lexema:   string(char),
					})
					nodes = append(nodes, &Node{
						NodeType: EOF,
						Lexema:   string("0"),
					})
					break lopp
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			if char == '=' {
				nodes = append(nodes, &Node{
					NodeType: LOWER_EQUALS,
					Lexema:   string(char),
				})
				continue
			}
			nodes = append(nodes, &Node{
				NodeType: LOWER,
				Lexema:   string(char),
			})
		case '>':
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					nodes = append(nodes, &Node{
						NodeType: GRATER,
						Lexema:   string(char),
					})
					nodes = append(nodes, &Node{
						NodeType: EOF,
						Lexema:   string("0"),
					})
					break lopp
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			if char == '=' {
				nodes = append(nodes, &Node{
					NodeType: GRATER_EQUALS,
					Lexema:   string(char),
				})
				continue
			}
			nodes = append(nodes, &Node{
				NodeType: GRATER,
				Lexema:   string(char),
			})
		case '!':
			char, _, err := t.buffer.ReadRune()
			if err != nil {
				if errors.Is(err, io.EOF) {
					nodes = append(nodes, &Node{
						NodeType: ERROR,
						Lexema:   "expected character = after !",
					})
					nodes = append(nodes, &Node{
						NodeType: EOF,
						Lexema:   string("0"),
					})
					break lopp
				}
				return nil, fmt.Errorf("error reading the rune: %w", err)
			}
			if char == '=' {
				nodes = append(nodes, &Node{
					NodeType: NOT_EQUALS,
					Lexema:   string(char),
				})
				continue
			}
			nodes = append(nodes, &Node{
				NodeType: ERROR,
				Lexema:   "expected character =",
			})
			break lopp
		}
	}
	return nodes, nil
}
