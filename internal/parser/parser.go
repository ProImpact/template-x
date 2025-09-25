package parser

import (
	"log"

	"github.com/ProImpact/templatex/internal/ast"
	"github.com/ProImpact/templatex/internal/lexer"
)

type Parser struct {
	lex   *lexer.Lexer
	Nodes []ast.Node
}

func NewParser(filePath string, addInclude bool) *Parser {
	lex, err := lexer.NewLexer(filePath, addInclude)
	if err != nil {
		log.Fatal(err)
	}
	return &Parser{
		lex: lex,
	}
}

func (p *Parser) Parse() {
	for {
		p.lex.Advance()
		tok := p.lex.Current()
		if tok.NodeType == lexer.EOF {
			return
		}
		if tok.NodeType == lexer.LBRACE {
			p.lex.Advance()
			tok = p.lex.Current()
			if tok.NodeType == lexer.LBRACE {
				p.lex.Advance()
				tok = p.lex.Current()
				switch tok.NodeType {
				case lexer.SLASH:
					comment, err := p.parseComment()
					if err != nil {
						log.Fatal(err)
					}
					p.Nodes = append(p.Nodes, comment)
				default:
					continue
				}
			} else {
				continue
			}
		} else {
			continue
		}
	}
}
