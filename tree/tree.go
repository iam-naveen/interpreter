package tree

type Node interface {
	String() string
}

type Stmt interface {
	Node
	Stmt()
	print(level int, prefix, out string, prev bool) string
}

type Expr interface {
	Node
	Expr()
	print(level int, prefix, out string, last bool) string
}

const (
	indent = "   "
	pipe = "│"
	Tee = "├──"
	Last = "└──"
	line = "─"
)
