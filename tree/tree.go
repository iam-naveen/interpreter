package tree

import "fmt"

// Basic interface for all nodes in the AST

type Node interface {
	String() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

// Program is the root node of the AST

type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out string
	for _, s := range p.Statements {
		out += fmt.Sprintf("%v", s)
	}
	return out
}
