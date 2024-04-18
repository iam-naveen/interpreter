package tree

import "fmt"


type Block struct {
	Children []Stmt
}
func (b *Block) String() string {
	var out string
	for _, s := range b.Children {
		out += fmt.Sprintf("%v", s)
	}
	return out
}
func (b *Block) Stmt() {}


type ExpressionStmt struct {
	Expression Expr
}

func (e *ExpressionStmt) String() string {
	return fmt.Sprintf("%v\n", e.Expression)
}
func (e *ExpressionStmt) StatementNode() {}


type Declaration struct {
	Datatype string
	Name     string
	Value    Expr
}
func (v *Declaration) String() string {
	return fmt.Sprintf("%s %s %s\n", v.Datatype, v.Name, v.Value)
}
func (v *Declaration) Stmt() {}


