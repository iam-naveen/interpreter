package lexer

import (
	"strings"
	"unicode/utf8"
)

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
