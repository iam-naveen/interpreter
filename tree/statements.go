package tree

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
)

// =====================================
// ======== PROGRAM STRUCTURE ==========
// =====================================
type Program struct {
	Children []Stmt
}

func (b *Program) String() string {
	var out string
	for _, s := range b.Children {
		out += fmt.Sprintf("%v", s)
	}
	return out
}

func (b *Program) Stmt() {}

func (s Program) Print(level int, prefix, out string) string {
	if len(s.Children) == 0 {
		return out
	}
	for i, stmt := range s.Children {
		if i == len(s.Children)-1 {
			out += stmt.print(level, prefix+Last, "", true)
		} else {
			out += stmt.print(level, prefix+Tee, "", false)
		}
	}
	return out
}
// =====================================
// ========= BLOCK STATEMENT ===========
// =====================================

type Block struct {
	Piece    lexer.Piece
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

func (b *Block) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, "{}")
	if len(b.Children) == 0 {
		return out
	}
	for i, stmt := range b.Children {
		if i == len(b.Children)-1 {
			out += stmt.print(level, prefix+Last, "", true)
		} else {
			out += stmt.print(level, prefix+Tee, "", false)
		}
	}
	return out
}

// =====================================
// ======== EXPRESSION STATEMENT =======
// =====================================

type ExpressionStmt struct {
	Expression Expr
}

func (e *ExpressionStmt) String() string {
	return fmt.Sprintf("%v\n", e.Expression)
}

func (e *ExpressionStmt) Stmt() {}

func (s *ExpressionStmt) print(level int, prefix, out string, last bool) string {
	out += s.Expression.print(level, prefix, "", last)
	return out
}

// =====================================
// ======= DECLARATION STATEMENT =======
// =====================================

type Declaration struct {
	Datatype lexer.Piece
	Name     lexer.Piece
	Value    Expr
}

func (v *Declaration) String() string {
	return fmt.Sprintf("%s %s %s\n", v.Datatype.Value, v.Name.Value, v.Value)
}
func (v *Declaration) Stmt() {}


func (s *Declaration) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Name.Value)
	margin := strings.Repeat(indent, level+1)
	out += s.Value.print(level+1, Last, pipe+margin, true)
	return out
}

func (s *Identifier) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

func (s *Number) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

func (s *StringLiteral) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s)
	return out
}

func (b *Boolean) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, b)
	return out
}

func (s *Binary) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Operator.Value)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Left.print(level+1, Tee, margin, last)
	out += s.Right.print(level+1, Last, margin, false)
	return out
}

func (s *Prefix) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Operator.Value)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Right.print(level+1, Last, margin, true)
	return out
}
