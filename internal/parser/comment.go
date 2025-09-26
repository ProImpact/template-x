package parser

import (
	"fmt"

	"github.com/ProImpact/templatex/internal/ast"
	"github.com/ProImpact/templatex/internal/lexer"
)

func (p *Parser) parseComment() (*ast.Comment, error) {
	tok := p.lex.Current()
	if tok.NodeType != lexer.SLASH {
		return nil, fmt.Errorf("expected \\ found a node %+v", tok)
	}
	p.lex.Advance()
	tok = p.lex.Current()
	if tok.NodeType != lexer.ASTERISIC {
		return nil, fmt.Errorf("expected * found a node %+v", tok)
	}
	node := new(ast.Comment)
	node.BaseNode = ast.BaseNode{
		NType: ast.NodeComment,
	}
	p.lex.Advance()
	tok = p.lex.Current()
	for tok.NodeType != lexer.ASTERISIC {
		if tok.NodeType == lexer.INDENTIFIER || tok.NodeType == lexer.KEYWORD {
			node.Words = append(node.Words, ast.NewWordNode(tok.Lexema))
			p.lex.Advance()
			tok = p.lex.Current()
			continue
		}
		return nil, fmt.Errorf("expected a IDENTIFIER or a KEYWORD found a node %+v", tok)
	}
	if tok.NodeType != lexer.ASTERISIC {
		return nil, fmt.Errorf("expected * found a node %+v", tok)
	}
	p.lex.Advance()
	tok = p.lex.Current()
	if tok.NodeType != lexer.SLASH {
		return nil, fmt.Errorf("expected \\ found a node %+v", tok)
	}
	return node, nil
}
