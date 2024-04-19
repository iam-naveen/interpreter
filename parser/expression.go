package parser

import (
	"fmt"
	"strconv"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

func (p *Parser) parseExpression(pre precedence) tree.Expr {
	prefix, ok := prefixHandlers[p.piece.Kind]
	if !ok {
		panic("No prefix handler for " + p.piece.Value)
	}
	left := prefix(p)
	// fmt.Println("Parsed", left)
	for p.piece.Kind != lexer.Eol && pre < precLookup[p.piece.Kind] {
		infix, has := infixHandlers[p.piece.Kind]
		if !has {
			return left
		}
		left = infix(p, left, precLookup[p.piece.Kind])
	}
	return left
}

func parseIdentifier(p *Parser) tree.Expr {
	variable := &tree.Identifier{
		Piece: *p.piece,
		Name: p.piece.Value,
	}
	p.move()
	return variable
}

func parseNumber(p *Parser) tree.Expr {
	number := &tree.Number{ Piece: *p.piece, }
	val, err := strconv.ParseInt(p.piece.Value, 10, 64)
	if err != nil {
		panic(err)
	}
	number.Value = val
	p.move()
	return number
}

func parseString(p *Parser) tree.Expr {
	stringLiteral := &tree.StringLiteral{
		Piece: *p.piece,
		Value: p.piece.Value,
	}
	p.move()
	return stringLiteral
}

func parseBoolean(p *Parser) tree.Expr {
	boolean := &tree.Boolean{
		Piece: *p.piece,
	}
	switch p.piece.Value {
	case "aam":
		boolean.Value = true
	case "illai":
		boolean.Value = false
	default:
		panic("Invalid boolean value")
	}
	p.move()
	return boolean
}

func parseGrouped(p *Parser) tree.Expr {
	// fmt.Println("Parsing group", p)
	p.move()
	expr := p.parseExpression(LOWEST)
	fmt.Println("Parsed group", expr)
	if p.piece.Kind != lexer.ParanClose {
		panic("Expected closing paranthesis")
	}
	p.move()
	return expr
}

func parsePrefix(p *Parser) tree.Expr {
	operator := p.piece
	p.move()
	right := p.parseExpression(UNARY)
	return &tree.Prefix{
		Operator: *operator,
		Right:    right,
	}
}

func parseInfix(p *Parser, left tree.Expr, bp precedence) tree.Expr {
	operator := p.piece
	p.move()
	right := p.parseExpression(bp)
	return &tree.Binary{
		Left:     left,
		Operator: *operator,
		Right:    right,
	}
}
