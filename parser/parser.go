package parser

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

// LookUp table for the keywords
// TODO: this is in two places now. Need to move it to a common place
var lookup = map[string]lexer.PieceType{
	"yen":   lexer.Keyword,
	"sol":   lexer.Keyword,
	"sollu": lexer.Keyword,
}

// Precedence levels for the operators
const (
	_ int = iota
	LOWEST
	EQUALS  // ==
	COMPARE // > or <
	PLUS    // +
	STAR    // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
)

type Parser struct {
	cur     *lexer.Piece
	prev    *lexer.Piece
	channel chan lexer.Piece
}

func (p Parser) String() string {
	C := strings.Trim(p.cur.Value, "\n")
	if p.prev == nil {
		return fmt.Sprintf("nil, %s", C)
	}
	P := strings.Trim(p.prev.Value, "\n")
	return fmt.Sprintf("%s, %s", P, C)
}

type state func(*Parser, *tree.Program) state

func RunParser(channel chan lexer.Piece) {
	parser := &Parser{channel: channel}
	program := &tree.Program{
		Statements: []tree.Statement{},
	}
	for state := startState; state != nil; {
		state = state(parser, program)
	}
	fmt.Println(program)
}

func (p *Parser) move() {
	for {
		select {
		case piece := <-p.channel:
			p.prev = p.cur
			p.cur = &piece
			fmt.Println("moving -->", p)
			return
		}

	}
}

func startState(parser *Parser, program *tree.Program) state {

	parser.move()

	if parser.cur.Kind == lexer.Keyword {
		return declarationState
	}

	// if the line is starting with an identifier,
	// then it is an Assignment statement
	if parser.cur.Kind == lexer.Identifier {
		return nil
	}

	// Kill the parser if we reach the end of the file
	if parser.cur.Kind == lexer.Eof {
		return nil
	}

	return nil
}

type Variable struct {
	Datatype string
	Name     string
	Value    string
}

// String implements tree.Statement.
func (v *Variable) String() string {
	return fmt.Sprintf("%s %s %s\n", v.Datatype, v.Name, v.Value)
}

// statementNode implements tree.Statement.
func (v *Variable) StatementNode() {}

func declarationState(parser *Parser, program *tree.Program) state {
	switch parser.cur.Value {
	case "yen":
		return createNumberVariable
	case "sol":
		return createStringVariable
	default:
		return nil
	}
}

func createNumberVariable(p *Parser, program *tree.Program) state {
	variable := &Variable{Datatype: "number"}
	if p.move(); p.cur.Kind == lexer.Identifier {
		variable.Name = p.cur.Value
	}
	if p.move(); p.cur.Kind == lexer.AssignmentOperator {
		program.Statements = append(program.Statements, variable)
		return expressionState
	}
	return startState
}

func expressionState(p *Parser, program *tree.Program) state {
	
	for p.move(); p.cur.Kind != lexer.Eol; p.move() {
		fmt.Println("expressionState")
	}

	return startState
}

func createStringVariable(p *Parser, program *tree.Program) state {
	fmt.Println("string variable created")
	return startState
}
