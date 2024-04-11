package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type PieceType int

const (
	Eof PieceType = iota
	Eol
	Keyword
	Identifier
	Number
	StringLiteral
	ArithmeticOperator
	LogicalOperator
	AssignmentOperator
	Unknown
)

var lookup = map[string]PieceType{
	"yen":    Keyword,
	"sol":    Keyword,
	"sollu":  Keyword,
}

type Piece struct {
	Kind  PieceType
	Value string
}

func (p Piece) String() string {
	switch p.Kind {
	case Keyword:
		return fmt.Sprintf("keyword: %s", p.Value)
	case Identifier:
		return fmt.Sprintf("identifier: %s", p.Value)
	case Number:
		return fmt.Sprintf("number: %s", p.Value)
	case StringLiteral:
		return fmt.Sprintf("string: %s", p.Value)
	case ArithmeticOperator:
		return fmt.Sprintf("arithmetic: %s", p.Value)
	case LogicalOperator:
		return fmt.Sprintf("logical: %s", p.Value)
	case AssignmentOperator:
		return fmt.Sprintf("assignment: %s", p.Value)
	case Eol:
		return "nextLine"
	case Eof:
		return "END"
	default:
		return fmt.Sprintf("unknown: %s", p.Value)
	}
}

type Lexer struct {
	input   []byte
	start   int // start position of the piece
	cur     int // current position in input
	size    int // size of the current piece
	channel chan Piece
}

func CreateLexer(input []byte) (*Lexer, chan Piece) {
	lex := &Lexer{
		input:   input,
		channel: make(chan Piece),
	}
	go lex.run()
	return lex, lex.channel
}

func (lex *Lexer) send(p PieceType) {
	piece := Piece{p, string(lex.input[lex.start:lex.cur])}
	lex.channel <- piece
	lex.start = lex.cur
}

func (lex *Lexer) next() (r rune) {
	if lex.cur >= len(lex.input) {
		lex.size = 0
		return rune(Eof)
	}
	r, lex.size = utf8.DecodeRuneInString(string(lex.input[lex.cur:]))
	lex.cur += lex.size
	return r
}

func (lex *Lexer) goBack() {
	lex.cur -= lex.size
}

func (lex *Lexer) peek() rune {
	r := lex.next()
	lex.goBack()
	return r
}

func (lex *Lexer) ignore() {
	lex.start = lex.cur
}

func (lex *Lexer) takeOne(valid string) bool {
	if strings.IndexRune(valid, lex.next()) >= 0 {
		return true
	}
	lex.goBack()
	return false
}

func (lex *Lexer) takeMany(valid string) {
	for strings.IndexRune(valid, lex.next()) >= 0 {
	}
	lex.goBack()
}

func (lex *Lexer) run() {
	for consumer := initial; consumer != nil; {
		consumer = consumer(lex)
	}
	close(lex.channel)
}

func consume(lex *Lexer) consumer {
	if unicode.IsLetter(rune(lex.input[lex.start])) {
		return consumeAlphaNumeric
	}
	if unicode.IsDigit(rune(lex.input[lex.start])) {
		return consumeNumber
	}
	return initial
}

func consumeAlphaNumeric(lex *Lexer) consumer {
	lex.takeMany("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")
	word := string(lex.input[lex.start:lex.cur])
	if kind, ok := lookup[word]; ok {
		lex.send(kind)
	} else {
		lex.send(Identifier)
	}
	return initial
}

func consumeNumber(lex *Lexer) consumer {
	lex.takeMany("0123456789")
	lex.send(Number)
	return initial
}

type consumer func(*Lexer) consumer

func initial(lex *Lexer) consumer {
	for {
		if lex.cur >= len(lex.input) {
			lex.send(Eof)
			return nil
		}
		if lex.takeOne("\n") {
			lex.send(Eol)
			continue
		}
		if lex.takeOne(" \t") {
			lex.ignore()
			continue
		}
		if lex.takeOne("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_") {
			return consumeAlphaNumeric
		}
		if lex.takeOne("0123456789") {
			return consumeNumber
		}
		if lex.takeOne("+*-/%") {
			lex.send(ArithmeticOperator)
			continue
		}
		if lex.takeOne("="){
			lex.send(AssignmentOperator)
			continue
		} else {
			lex.next()
			lex.send(Unknown)
		}
	}
}
