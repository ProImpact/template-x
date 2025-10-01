package ast

// AstNode used to parse complex tree structures
// Eg: if -> if ( .Name = "" and (.Marco[0]["name"] == 21))
type AstNode struct {
	Nodes []*AstNode
}
