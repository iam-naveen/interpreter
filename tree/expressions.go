package tree

import (
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
)

type Identifier struct {
	Piece lexer.Piece
	Name string
}
func (i *Identifier) String() string {
	return i.Name
}
func (i *Identifier) Expr() {}



type Number struct {
	Piece lexer.Piece
	Value int64
}
func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}
func (n *Number) Expr() {}



type StringLiteral struct {
	Piece lexer.Piece
	Value string
}
func (s *StringLiteral) String() string {
	return s.Value
}
func (s *StringLiteral) Expr() {}



type Boolean struct {
	Piece lexer.Piece
	Value bool
}
func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}
func (b *Boolean) Expr() {}


type Array struct {
	Piece lexer.Piece
	Elements []Expr
}
func (a *Array) String() string {
	return fmt.Sprintf("%v", a.Elements)
}
func (a *Array) Expr() {}



// ===========================
// === COMPLEX EXPRESSIONS ===
// ===========================

type Binary struct {
	Left     Expr
	Operator lexer.Piece
	Right    Expr
}
func (b *Binary) String() string {
	return fmt.Sprintf("(%v %s %v)", b.Left, b.Operator.Value, b.Right)
}
func (b *Binary) Expr() {}


type Prefix struct {
	Operator lexer.Piece
	Right    Expr
}
func (p *Prefix) String() string {
	return fmt.Sprintf("(%s%v)", p.Operator.Value, p.Right)
}
func (p *Prefix) Expr() {}
