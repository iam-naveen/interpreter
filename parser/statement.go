package parser

import (
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

func parseStatement(p *Parser) tree.Stmt {
	switch p.piece.Kind {
	case lexer.DataType:
		return p.parseDeclarationStatement()
	default:
		return p.parserExpressionStatement()
	}
}

func (p *Parser) parseDeclarationStatement() tree.Stmt {
	stmt := &tree.Declaration{ Datatype: *p.piece }
	p.move()
	stmt.Name = *p.piece
	p.move()
	if p.piece.Kind == lexer.Assign {
		p.move()
		stmt.Value = p.parseExpression(LOWEST)
		if p.logEnabled {
			fmt.Println("Declaration statement parsed\n\t", stmt);
		}
		p.move()
		return stmt
	}
	return nil
}

func (p *Parser) parserExpressionStatement() tree.Stmt {
	expr := p.parseExpression(LOWEST)
	switch p.piece.Kind {
	case lexer.If:
		return p.parseIfStatement(expr)
	case lexer.Print:
		printStmt := &tree.PrintStmt{ Piece: *p.piece }
		printStmt.Value = expr
		p.move()
		if p.piece.Kind != lexer.Eol {
			panic("Expected ;")
		}
		p.move()
		return printStmt
	default:
		p.move()
		return &tree.ExpressionStmt{ Expression: expr }
	}
}

func (p *Parser) parseIfStatement(expr tree.Expr) tree.Stmt {
	ifStmt := &tree.IfStmt{ Condition: expr }
	p.move()
	if p.piece.Kind != lexer.BraceOpen {
		panic("Expected '{'")
	}
	ifStmt.Then = p.parseBlockStatement()
	if p.piece.Kind == lexer.Else {
		p.move()
		if p.piece.Kind != lexer.BraceOpen {
			panic("Expected '{'")
		}
		ifStmt.Else = p.parseBlockStatement()
	}
	return ifStmt
}

func (p *Parser) parseBlockStatement() *tree.Block {
	block := &tree.Block{ Piece: *p.piece }
	p.move()
	for p.piece.Kind != lexer.BraceClose {
		stmt := parseStatement(p)
		if stmt != nil {
			block.Children = append(block.Children, stmt)
		}
	}
	if p.piece.Kind != lexer.BraceClose {
		panic("Expected '}'")
	}
	p.move()
	return block
}
