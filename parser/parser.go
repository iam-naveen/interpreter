package parser

import (
	"fmt"
	"strings"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

type Parser struct {
	piece     *lexer.Piece
	prev    *lexer.Piece
	channel chan lexer.Piece
	tree []tree.Stmt
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
	createStates()
	parser := &Parser{
		channel: channel,
		tree: []tree.Stmt{},
	}
	parser.move()
	for handler, ok := stmtHandlers[parser.piece.Kind]; ok; {
		stmt := handler(parser)
		if stmt != nil {
			parser.tree = append(parser.tree, stmt)
		}
	}
}
