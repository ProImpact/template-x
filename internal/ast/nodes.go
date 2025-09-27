package ast

import (
	"encoding/json"
	"fmt"
)

type NodeType int

const (
	NodeProgram NodeType = iota
	NodeVariable
	NodeDeclaration
	NodeAssigment
	NodeParameter
	NodeFunctionCall
	NodeWord   // word
	NodeString // "word"
	NodeInteger
	NodeFloat
	NodeChar
	NodeIf
	NodeElse
	NodeElseIf
	NodeRange
	NodeComparison
	NodeDefine
	NodeTemplate
	NodeBoolExpr
	NodePipe
	NodeComment
	NodeExpression
	NodeBoolExpression
	NodeArrayAccess
	NodeMapAccess
)

type Node interface {
	Type() NodeType
	String() string
	Accept(visitor Visitor) any
	json.Marshaler
}

type BaseNode struct {
	NType NodeType `json:"n_type,omitempty"`
}

func (b *BaseNode) Type() NodeType {
	return b.NType
}

type Program struct {
	BaseNode     `json:"base_node,omitempty"`
	Declarations []Node `json:"declarations,omitempty"`
}

func NewProgram() *Program {
	return &Program{
		BaseNode: BaseNode{
			NType: NodeProgram,
		},
		Declarations: make([]Node, 0),
	}
}

func (p *Program) String() string {
	return fmt.Sprintf("Program: {declarations: %d}", len(p.Declarations))
}

type Comment struct {
	BaseNode `json:"base_node,omitempty"`
	Words    []Node `json:"words,omitempty"`
}

func (c *Comment) Accept(visitor Visitor) any {
	for _, word := range c.Words {
		fmt.Println(word)
	}
	return nil
}

func (c *Comment) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"node_type": "Comment",
		"words":     c.Words,
	}
	return json.Marshal(&data)
}

func (c *Comment) String() string {
	return "CommentNode"
}

type Word struct {
	BaseNode
	Value string `json:"value,omitempty"`
}

func (w *Word) String() string {
	return w.Value
}

func NewWordNode(value string) *Word {
	return &Word{
		BaseNode: BaseNode{
			NType: NodeWord,
		},
		Value: value,
	}
}

func (w *Word) Accept(visitor Visitor) any {
	return w.String()
}

func (c *Word) MarshalJSON() ([]byte, error) {
	data := map[string]any{
		"node_type": "Word",
		"lexema":    c.Value,
	}
	return json.Marshal(&data)
}
