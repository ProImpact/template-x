package ast

import "github.com/ProImpact/templatex/internal/lexer"

type Node interface {
	Start() lexer.Position
	End() lexer.Position
	String() string
}

type Expresion interface {
	Node
	expresionNode()
}

type Statement interface {
	Node
	statementNode()
}

type CommentNode struct {
	Words         []Node
	StartPosition int
	EndPosition   int
}

func (c *CommentNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *CommentNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *CommentNode) String() string {
	return "CommentNode"
}

func NewCommentNode(start, end int, words []Node) Node {
	return &CommentNode{
		Words:         words,
		StartPosition: start,
		EndPosition:   end,
	}
}

type WordComment struct {
	Lexema        string
	StartPosition int
	EndPosition   int
}

func (c *WordComment) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *WordComment) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *WordComment) String() string {
	return "WordCommentNode"
}

func NewWordCommentNode(start, end int, lexema string) Node {
	return &WordComment{
		Lexema:        lexema,
		StartPosition: start,
		EndPosition:   end,
	}
}

type TemplateDefinitionNode struct {
	TemplateName  string
	StartPosition int
	EndPosition   int
	Args          *Node
}

func (c *TemplateDefinitionNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *TemplateDefinitionNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *TemplateDefinitionNode) String() string {
	return "TemplateDefinitionNode"
}

func NewTemplateDefinitionNode(start, end int, templateName string, args *Node) Node {
	return &TemplateDefinitionNode{
		TemplateName:  templateName,
		StartPosition: start,
		EndPosition:   end,
		Args:          args,
	}
}

type StatementFinalizationNode struct {
	StartPosition int
	EndPosition   int
}

func (c *StatementFinalizationNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *StatementFinalizationNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *StatementFinalizationNode) String() string {
	return "StatementFinalizationNode"
}

func NewStatementFinalization(start, end int) Node {
	return &StatementFinalizationNode{
		StartPosition: start,
		EndPosition:   end,
	}
}

type FieldNode struct {
	Lexema        string
	StartPosition int
	EndPosition   int
}

func (c *FieldNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *FieldNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *FieldNode) String() string {
	return "FieldNode"
}

func NewFieldNode(start, end int, lexema string) Node {
	return &FieldNode{
		Lexema:        lexema,
		StartPosition: start,
		EndPosition:   end,
	}
}

type ArrayAccessNode struct {
	Index         int
	StartPosition int
	EndPosition   int
}

func (c *ArrayAccessNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *ArrayAccessNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *ArrayAccessNode) String() string {
	return "ArrayAccessNode"
}

func NewArrayAccessNode(start, end int, index int) Node {
	return &ArrayAccessNode{
		Index:         index,
		StartPosition: start,
		EndPosition:   end,
	}
}

type MapAccessNode struct {
	Key           string
	StartPosition int
	EndPosition   int
}

func (c *MapAccessNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *MapAccessNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *MapAccessNode) String() string {
	return "MapAccessNode"
}

func NewMapAccessNode(start, end int, key string) Node {
	return &MapAccessNode{
		Key:           key,
		StartPosition: start,
		EndPosition:   end,
	}
}

type TemplateBlockNode struct {
	TemplateName  string
	StartPosition int
	EndPosition   int
	Args          *Node
}

func (c *TemplateBlockNode) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *TemplateBlockNode) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *TemplateBlockNode) String() string {
	return "TemplateBlockNode"
}

func NewTemplateBlockNode(start, end int, templateName string, args *Node) Node {
	return &TemplateBlockNode{
		TemplateName:  templateName,
		StartPosition: start,
		EndPosition:   end,
		Args:          args,
	}
}

type Literal struct {
	Lexema        string
	LiteralType   lexer.TokenType
	StartPosition int
	EndPosition   int
}

func (c *Literal) End() lexer.Position {
	return lexer.Position(c.EndPosition)
}

func (c *Literal) Start() lexer.Position {
	return lexer.Position(c.StartPosition)
}

func (c *Literal) String() string {
	switch c.LiteralType {
	case lexer.TEXT:
		return "LiteralString"
	case lexer.FLOAT:
		return "LiteralFloat"
	case lexer.NUMBER:
		return "LiteralNumber"
	case lexer.FALSE:
		return "LiteralFalse"
	case lexer.TRUE:
		return "LiteralTrue"
	default:
		return "Invalid"
	}
}

func NewLiteral(start, end int, tokenType lexer.TokenType, lexema string) Node {
	return &Literal{
		Lexema:        lexema,
		StartPosition: start,
		EndPosition:   end,
		LiteralType:   tokenType,
	}
}
