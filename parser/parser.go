package parser

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

type Parser struct {
	piece      *lexer.Piece
	prev       *lexer.Piece
	channel    chan lexer.Piece
	logEnabled bool
}

func (p Parser) String() string {
	C := strings.Replace(p.piece.Value, "\n", "eol", 1)
	if p.prev == nil {
		return fmt.Sprintf("nil, %s", C)
	}
	P := strings.Replace(p.prev.Value, "\n", "eol", 1)
	return fmt.Sprintf("%s, %s", P, C)
}

func Parse(channel chan lexer.Piece, logging bool) *tree.Program {
	parser := &Parser{channel: channel, logEnabled: logging}
	parser.createHandlers()
	parser.move()

	program := &tree.Program{}

	handler, present := stmtHandlers[parser.piece.Kind]
	for present {
		stmt := handler(parser)
		if stmt != nil {
			program.Children = append(program.Children, stmt)
		} else {
			fmt.Println("Error parsing statement", parser)
			break
		}
		if parser.piece.Kind == lexer.Eof {
			break
		}
		handler, present = stmtHandlers[parser.piece.Kind]
	}
	return program
}
