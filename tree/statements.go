package tree

import (
	"bytes"
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
	margin := strings.Repeat(pipe+indent, level+1)
	for i, stmt := range b.Children {
		if i == len(b.Children)-1 {
			out += stmt.print(level+1, Last, margin, true)
		} else {
			out += stmt.print(level+1, Tee, margin, false)
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
	Datatype string
	Name     lexer.Piece
	Value    Expr
}

func (v *Declaration) String() string {
	return fmt.Sprintf("%s %s %s\n", v.Datatype, v.Name.Value, v.Value)
}
func (v *Declaration) Stmt() {}

func (s *Declaration) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, s.Name.Value)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Value.print(level+1, Last, margin, true)
	return out
}

// =====================================
// ========= IF STATEMENT ==============
// =====================================

type IfStmt struct {
	Piece     lexer.Piece
	Condition Expr
	Then      *Block
	Else      Stmt
}

func (i *IfStmt) String() string {
	out := fmt.Sprintf("if %v %v\n", i.Condition, i.Then)
	if i.Else != nil {
		out += fmt.Sprintf("else %v", i.Else)
	}
	return out
}

func (i *IfStmt) Stmt() {}

func (s *IfStmt) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s if %s\n", prefix, s.Condition)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Then.print(level+1, Tee, margin, false)
	if s.Else != nil {
		out += s.Else.print(level+1, prefix, margin, true)
	}
	return out
}

// =====================================
// ======== WHILE STATEMENT ============
// =====================================

type WhileStmt struct {
	Piece     lexer.Piece
	Condition Expr
	Body      *Block
}

func (w *WhileStmt) String() string {
	return fmt.Sprintf("while %v %v\n", w.Condition, w.Body)
}

func (w *WhileStmt) Stmt() {}

func (s *WhileStmt) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s while %s\n", prefix, s.Condition)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Body.print(level+1, Tee, margin, true)
	return out
}

// =====================================
// ======== FOR STATEMENT ==============
// =====================================

type ForStmt struct {
	Piece lexer.Piece
	Count Expr
	Body  *Block
}

func (f *ForStmt) String() string {
	return fmt.Sprintf("for %v %v\n", f.Count, f.Body)
}

func (f *ForStmt) Stmt() {}

func (s *ForStmt) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s times\n", prefix, s.Count)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Body.print(level+1, Tee, margin, true)
	return out
}

// =====================================
// ======== PRINT STATEMENT ============
// =====================================

type PrintStmt struct {
	Piece lexer.Piece
	Value Expr
}

func (p *PrintStmt) String() string {
	return fmt.Sprintf("print %v\n", p.Value)
}

func (p *PrintStmt) Stmt() {}

func (s *PrintStmt) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s %s\n", prefix, "print")
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Value.print(level+1, Last, margin, true)
	return out
}

// =====================================
// ======== FUNCTION ===================
// =====================================

type Function struct {
	Piece  lexer.Piece
	Name   lexer.Piece
	Args   []Expr
	Return lexer.Piece
	Body   *Block
}

func (f *Function) String() string {
	out := bytes.Buffer{}
	out.WriteString(fmt.Sprintf("fn %s(", f.Name.Value))
	for i, arg := range f.Args {
		out.WriteString(arg.String())
		if i < len(f.Args)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(") %v\n")
	out.WriteString(f.Body.String())
	return out.String()
}

func (f *Function) Stmt() {}

func (s *Function) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s fn %s\n", prefix, s.Name.Value)
	return out
}

// =====================================
// ======== RETURN STATEMENT ===========
// =====================================

type ReturnStmt struct {
	Piece lexer.Piece
	Value Expr
}

func (r *ReturnStmt) String() string {
	return fmt.Sprintf("return %v\n", r.Value)
}

func (r *ReturnStmt) Stmt() {}

func (s *ReturnStmt) print(level int, prefix, out string, last bool) string {
	out += fmt.Sprintf("%s return\n", prefix)
	margin := strings.Repeat(pipe+indent, level+1)
	out += s.Value.print(level+1, Last, margin, true)
	return out
}
