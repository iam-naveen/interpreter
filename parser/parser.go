package parser

import (
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
)

// LookUp table for the keywords
// TODO: this is in two places now. Need to move it to a common place
var lookup = map[string]lexer.PieceType{
	"yen":   lexer.Keyword,
	"sol":   lexer.Keyword,
	"sollu": lexer.Keyword,
}

type Parser struct {
	cur     *lexer.Piece
	prev    *lexer.Piece
	channel chan lexer.Piece
}

func (p *Parser) String() string {
	return fmt.Sprintf("prev: %s, cur: %s", p.prev.Value, p.cur.Value)
}

type state func(*Parser) state

func RunParser(channel chan lexer.Piece) {
	parser := &Parser{channel: channel}
	for state := start; state != nil; {
		state = state(parser)
	}
}

func (p *Parser) move(piece *lexer.Piece) {
	p.prev = p.cur
	p.cur = piece
}

func start(p *Parser) state {
	for {
		select {
		case piece := <-p.channel:
			p.move(&piece)

			// if the line is starting with a keyword,
			// then it is a Declaration statement
			if p.cur.Kind == lexer.Keyword {
				return createVariable
			}

			// if the line is starting with an identifier,
			// then it is an Assignment statement
			if p.cur.Kind == lexer.Identifier {
				return nil
			}

			// Kill the parser if we reach the end of the file
			if p.cur.Kind == lexer.Eof {
				return nil
			}
		}
	}
}

func createVariable(p *Parser) state {
	switch p.cur.Value {
	case "yen":
		fmt.Println("number variable")
		return createNumberVariable
	case "sol":
		fmt.Println("string variable")
		return createStringVariable
	}
	return nil
}

func createNumberVariable(p *Parser) state {
	fmt.Println("number variable created")
	return start
}

func createStringVariable(p *Parser) state {
	fmt.Println("string variable created")
	return start
}
