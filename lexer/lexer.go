package lexer

import (
	"github.com/iam-naveen/compiler/token"
)

type Lexer struct {
	input []byte
	cur   int  // current position in input (points to current char)
	peek  int  // current reading position in input (after current char)
	char  byte // current char under examination
}

func New(input []byte) *Lexer {
	lex := &Lexer{input: input}
	return lex
}

func (lex *Lexer) readChar() {
	// set the char
	if lex.peek >= len(lex.input) {
		lex.char = 0 // mark as end - ASCII code for "NUL"
	} else {
		lex.char = lex.input[lex.peek]
	}
	// move the pointers
	lex.cur = lex.peek
	lex.peek += 1
}

func (lex *Lexer) NextToken() token.Token {

	tok := token.Token{}
	lex.readChar()

	// neglect whitespaces
	for lex.char == ' ' || lex.char == '\t' {
		lex.readChar()
	}

	if lex.char == 0 {
		tok.Value = ""
		tok.Type = token.EOF
	} else if lex.char == '\n' {
		tok.Value = ""
		tok.Type = token.EOL
	} else if lex.char == '=' {
		if lex.input[lex.peek] == '=' {
			lex.readChar()
			tok.Value = "=="
			tok.Type = token.EQ
		} else {
			tok.Value = "="
			tok.Type = token.ASSIGN
		}
	} else if lex.char == '+' {
		tok.Value = "+"
		tok.Type = token.PLUS
	} else if lex.char == '-' {
		tok.Value = "-"
		tok.Type = token.MINUS
	} else if lex.char == '!' {
		if lex.input[lex.peek] == '=' {
			lex.readChar()
			tok.Value = "!="
			tok.Type = token.NEQ
		} else {
			tok.Value = "!"
			tok.Type = token.NOT
		}
	} else if lex.char == '*' {
		tok.Value = "*"
		tok.Type = token.STAR
	} else if lex.char == '/' {
		tok.Value = "/"
		tok.Type = token.SLASH
	} else if lex.char == '<' {
		tok.Value = "<"
		tok.Type = token.LT
	} else if lex.char == '>' {
		tok.Value = ">"
		tok.Type = token.GT
	} else if lex.char == ';' {
		tok.Value = ";"
		tok.Type = token.SEMICOLON
	} else if lex.char == '(' {
		tok.Value = "("
		tok.Type = token.LPAREN
	} else if lex.char == ')' {
		tok.Value = ")"
		tok.Type = token.RPAREN
	} else if lex.char == '{' {
		tok.Value = "{"
		tok.Type = token.LBRACE
	} else if lex.char == '}' {
		tok.Value = "}"
		tok.Type = token.RBRACE
	} else if lex.char == ',' {
		tok.Value = ","
		tok.Type = token.COMMA
	} else if lex.char == '"' {
		tok.Value = lex.getString()
		tok.Type = token.STR
	} else {
		if isLetter(lex.char) {
			tok.Value = lex.getIdentifier()
			tok.Type = token.LookUp(tok.Value)
		} else if isDigit(lex.char) {
			tok.Value = lex.getNumber()
			tok.Type = token.INT
		}
	}
	return tok
}

func isLetter(char byte) bool {
	// define here the rules for variable names
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || (char == '_')
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (lex *Lexer) getString() string {
	start := lex.cur + 1
	for {
		lex.readChar()
		if lex.char == '"' || lex.char == 0 {
			break
		}
	}
	return string(lex.input[start:lex.cur])
}

func (lex *Lexer) getNumber() string {
	start := lex.cur
	for {
		if !lex.isNextDigit() {
			return string(lex.input[start:lex.peek])
		}
		lex.readChar()
	}
}

func (lex *Lexer) getIdentifier() string {
	start := lex.cur
	for {
		if !lex.isNextLetter() && !lex.isNextDigit() {
			return string(lex.input[start:lex.peek])
		}
		lex.readChar()
	}
}

func (lex *Lexer) isNextEnd() bool {
	return lex.peek < len(lex.input) && lex.input[lex.peek] == '\n'
}

func (lex *Lexer) isNextLetter() bool {
	return lex.peek < len(lex.input) && isLetter(lex.input[lex.peek])
}

func (lex *Lexer) isNextDigit() bool {
	return lex.peek < len(lex.input) && isDigit(lex.input[lex.peek])
}
