package parser

import (
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/tree"
)

func parseStatement(p *Parser) tree.Stmt {
	switch p.piece.Kind {
	case lexer.DataType:
		return p.parseDeclarationStatement()
	default:
		return p.parserExpressionStatement()
	}
}

func (p *Parser) parseDeclarationStatement() tree.Stmt {
	stmt := &tree.Declaration{ Datatype: *p.piece }
	p.move()
	stmt.Name = *p.piece
	p.move()
	if p.piece.Kind == lexer.Assign {
		p.move()
		stmt.Value = p.parseExpression(LOWEST)
		if p.logEnabled {
			fmt.Println("Declaration statement parsed\n\t", stmt);
		}
		p.move()
		return stmt
	}
	return nil
}

func (p *Parser) parserExpressionStatement() tree.Stmt {
	stmt := &tree.ExpressionStmt{
		Expression: p.parseExpression(LOWEST),
	}
	if p.logEnabled {
		fmt.Println("Expression statement parsed\n\t", stmt)
	}
	p.move()
	return stmt
}
