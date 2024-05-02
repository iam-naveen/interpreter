package parser

import (
	"fmt"
	"os"

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
	var datatype string
	var name lexer.Piece
	if p.piece.Value == "yen" {
		datatype = "INTEGER"
	} else if p.piece.Value == "sol" {
		datatype = "STRING"
	}
	p.move()
	if p.piece.Kind != lexer.Identifier {
		fmt.Println("Expected identifier got " + p.piece.Value)
		os.Exit(1)
	}
	name = *p.piece
	p.move()
	switch p.piece.Kind {
	case lexer.Eol:
		stmt := &tree.Declaration{
			Datatype: datatype,
			Name:     name,
		}
		if p.logEnabled {
			fmt.Println("Declaration statement parsed\n\t", stmt)
		}
		p.move()
		return stmt
	case lexer.Assign:
		p.move()
		stmt := &tree.Declaration{
			Datatype: datatype,
			Name:     name,
			Value:    p.parseExpression(LOWEST),
		}
		if p.logEnabled {
			fmt.Println("Declaration statement parsed\n\t", stmt)
		}
		p.move()
		return stmt

	// TODO: Implement builtin function calls
	// case lexer.Identifier:
	// 	if _, ok := evaluator.Builtins[p.piece.Value]; ok {
	// 		fmt.Println("Found builtin function", p.piece.Value)
	// 		p.move()
	// 	}
	// 	return nil

	case lexer.Input:
		inputStmt := &tree.Input{
			Piece: *p.piece,
			Variable: tree.Identifier{
				Piece: *p.prev,
				Name:  p.prev.Value,
			},
			DataType: datatype,
		}
		p.move()
		if p.piece.Kind != lexer.Eol {
			fmt.Println("Expected ;")
			os.Exit(1)
		}
		p.move()
		return inputStmt
	default:
		return nil
	}
}

func (p *Parser) parserExpressionStatement() tree.Stmt {
	expr := p.parseExpression(LOWEST) // After parsing, we are at the next token
	switch p.piece.Kind {
	case lexer.If:
		return p.parseIfStatement(expr)
	case lexer.While:
		return p.parseWhileStatement(expr)
	case lexer.For:
		return p.parseForStatement(expr)
	case lexer.Print:
		printStmt := &tree.PrintStmt{Piece: *p.piece}
		printStmt.Value = expr
		p.move()
		if p.piece.Kind != lexer.Eol {
			panic("Expected ;")
		}
		p.move()
		return printStmt
	default:
		p.move()
		return &tree.ExpressionStmt{Expression: expr}
	}
}

func (p *Parser) parseWhileStatement(expr tree.Expr) tree.Stmt {
	whileStmt := &tree.WhileStmt{Condition: expr}
	p.move()
	if p.piece.Kind != lexer.BraceOpen {
		fmt.Println("Expected '{' after 'varaikkum'")
		os.Exit(1)
	}
	whileStmt.Body = p.parseBlockStatement()
	return whileStmt
}

func (p *Parser) parseForStatement(expr tree.Expr) tree.Stmt {
	forStmt := &tree.ForStmt{Piece: *p.piece, Count: expr}
	p.move()
	if p.piece.Kind != lexer.BraceOpen {
		fmt.Println("Expected '{' after 'murai'")
		os.Exit(1)
	}
	forStmt.Body = p.parseBlockStatement()
	return forStmt
}

func (p *Parser) parseIfStatement(expr tree.Expr) tree.Stmt {
	ifStmt := &tree.IfStmt{Condition: expr}
	p.move()
	if p.piece.Kind != lexer.BraceOpen {
		panic("Expected '{'")
	}
	ifStmt.Then = p.parseBlockStatement()
	if p.piece.Kind == lexer.Else {
		p.move()
		if p.piece.Kind != lexer.BraceOpen {
			ifStmt.Else = p.parserExpressionStatement()
			return ifStmt
		}
		ifStmt.Else = p.parseBlockStatement()
	}
	return ifStmt
}

func (p *Parser) parseBlockStatement() *tree.Block {
	block := &tree.Block{Piece: *p.piece}
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
