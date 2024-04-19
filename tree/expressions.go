package tree

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
)

// ============================
// ======== IDENTIFIER  =======
// ============================

type Identifier struct {
	Piece lexer.Piece
	Name  string
}

func (i *Identifier) String() string {
	return i.Name
}

func (i *Identifier) Expr() {}

func (s *Identifier) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

// ============================
// ======== NUMBER  ===========
// ============================

type Number struct {
	Piece lexer.Piece
	Value int64
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

func (n *Number) Expr() {}

func (s *Number) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

// ============================
// ======== STRING  ===========
// ============================

type StringLiteral struct {
	Piece lexer.Piece
	Value string
}

func (s *StringLiteral) String() string {
	return s.Value
}

func (s *StringLiteral) Expr() {}

func (s *StringLiteral) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

// ============================
// ======== BOOLEAN ===========
// ============================

type Boolean struct {
	Piece lexer.Piece
	Value bool
}

func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}

func (b *Boolean) Expr() {}

func (b *Boolean) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, b)
	return out
}

// ============================
// ======== ARRAY ==============
// ============================

type Array struct {
	Piece    lexer.Piece
	Elements []Expr
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.Elements)
}

func (a *Array) Expr() {}

// ===========================
// === COMPLEX EXPRESSIONS ===
// ===========================

// ============================
// ======== BINARY =============
// ============================

type Binary struct {
	Left     Expr
	Operator lexer.Piece
	Right    Expr
}

func (b *Binary) String() string {
	return fmt.Sprintf("(%v %s %v)", b.Left, b.Operator.Value, b.Right)
}

func (b *Binary) Expr() {}

func (s *Binary) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Operator.Value)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Left.print(level+1, Tee, margin, last)
	out += s.Right.print(level+1, Last, margin, false)
	return out
}

// ============================
// ======== PREFIX =============
// ============================

type Prefix struct {
	Operator lexer.Piece
	Right    Expr
}

func (p *Prefix) String() string {
	return fmt.Sprintf("(%s%v)", p.Operator.Value, p.Right)
}

func (p *Prefix) Expr() {}

func (s *Prefix) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Operator.Value)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Right.print(level+1, Last, margin, true)
	return out
}
