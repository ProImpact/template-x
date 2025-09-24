package ast

type Visitor interface {
	VisitComment(*Comment) any
}
