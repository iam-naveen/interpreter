package tree

// Basic interface for all nodes in the AST

type Node interface {
	Position() int
	Literal() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST

type Program struct {
	Statements []Statement
}

func (p *Program) Position() int {
	if len(p.Statements) > 0 {
		return p.Statements[0].Position()
	}
	return 0
}

func (p *Program) Literal() string {
	return ""
}
