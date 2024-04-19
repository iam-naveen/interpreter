package parser

import (
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

type precedence int

const (
	LOWEST precedence = iota
	COMMA
	ASSIGNMENT
	LOGICAL
	RELATIONAL
	ADDITIVE
	MULTIPLICATIVE
	UNARY
	CALL
	MEMBER
	PRIMARY
)

type stmtHandler func(p *Parser) tree.Stmt
type prefixHandler func(p *Parser) tree.Expr
type infixHandler func(p *Parser, left tree.Expr, bp precedence) tree.Expr

type stmtLookup map[lexer.PieceType]stmtHandler
type prefixLookup map[lexer.PieceType]prefixHandler
type infixLookup map[lexer.PieceType]infixHandler
type precedenceLookup map[lexer.PieceType]precedence

var precLookup = precedenceLookup{}
var prefixHandlers = prefixLookup{}
var infixHandlers = infixLookup{}
var stmtHandlers = stmtLookup{}

func setInfixHandler(kind lexer.PieceType, bp precedence, handler infixHandler) {
	precLookup[kind] = bp
	infixHandlers[kind] = handler
}

func setPrefixHandler(kind lexer.PieceType, handler prefixHandler) {
	precLookup[kind] = PRIMARY
	prefixHandlers[kind] = handler
}

func setStmtHandler(kind lexer.PieceType, handler stmtHandler) {
	precLookup[kind] = LOWEST
	stmtHandlers[kind] = handler
}

func (p *Parser) createHandlers() {

	fmt.Println("Creating handlers...")
	fmt.Println()

	setStmtHandler(lexer.DataType, parseStatement)
	setStmtHandler(lexer.Identifier, parseStatement)
	setStmtHandler(lexer.Number, parseStatement)
	setStmtHandler(lexer.Boolean, parseStatement)

	setPrefixHandler(lexer.Identifier, parseIdentifier)
	setPrefixHandler(lexer.Number, parseNumber)
	setPrefixHandler(lexer.StringLiteral, parseString)
	setPrefixHandler(lexer.Boolean, parseBoolean)
	setPrefixHandler(lexer.Minus, parsePrefix)
	setPrefixHandler(lexer.Bang, parsePrefix)
	setPrefixHandler(lexer.ParanOpen, parseGrouped)
	// setPrefixHandler(lexer.BracketOpen, parseArray)

	setInfixHandler(lexer.Plus, ADDITIVE, parseInfix)
	setInfixHandler(lexer.Minus, ADDITIVE, parseInfix)
	setInfixHandler(lexer.Star, MULTIPLICATIVE, parseInfix)
	setInfixHandler(lexer.Slash, MULTIPLICATIVE, parseInfix)
	setInfixHandler(lexer.Percent, MULTIPLICATIVE, parseInfix)
	setInfixHandler(lexer.Equal, RELATIONAL, parseInfix)
	setInfixHandler(lexer.NotEqual, RELATIONAL, parseInfix)
	setInfixHandler(lexer.Assign, ASSIGNMENT, parseInfix)
}
