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

func infix(kind lexer.PieceType, bp precedence, handler infixHandler) {
	precLookup[kind] = bp
	infixHandlers[kind] = handler
}

func prefix(kind lexer.PieceType, handler prefixHandler) {
	precLookup[kind] = PRIMARY
	prefixHandlers[kind] = handler
}

func stmt(kind lexer.PieceType, handler stmtHandler) {
	precLookup[kind] = LOWEST
	stmtHandlers[kind] = handler
}

func createStates() {
	fmt.Println("Creating states yet to be implemented")
}
