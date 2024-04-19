package parser

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
	"github.com/sanity-io/litter"
)

type Parser struct {
	piece   *lexer.Piece
	prev    *lexer.Piece
	channel chan lexer.Piece
}

func (p Parser) String() string {
	C := strings.Replace(p.piece.Value, "\n", "eol", 1)
	if p.prev == nil {
		return fmt.Sprintf("nil, %s", C)
	}
	P := strings.Replace(p.prev.Value, "\n", "eol", 1)
	return fmt.Sprintf("%s, %s", P, C)
}

func Run(channel chan lexer.Piece) {
	parser := &Parser{ channel: channel }
	parser.createHandlers()
	parser.move()

	program := &tree.Block{}

	handler, present := stmtHandlers[parser.piece.Kind]
	for present {
		stmt := handler(parser)
		if stmt != nil {
			 program.Children = append(program.Children, stmt)
		} else {
			fmt.Println("Error parsing statement", parser)
			break
		}
		handler, present = stmtHandlers[parser.piece.Kind]
	}

	litter.Dump(program)
	fmt.Println(program.Print(0, "", ""))
}
