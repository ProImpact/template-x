package ast

import (
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
}

type BaseNode struct {
	NType NodeType
}

func (b *BaseNode) Type() NodeType {
	return b.NType
}

type Program struct {
	BaseNode
	Declarations []Node
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
	BaseNode
	Words []Node
}
