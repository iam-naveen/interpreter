package tree

type Node interface {
	String() string
}

type Stmt interface {
	Node
	Stmt()
}

type Expr interface {
	Node
	Expr()
}

