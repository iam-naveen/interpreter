package lexer

import "fmt"

type PieceType int

const (
	Eof PieceType = iota
	Eol
	Keyword
	Identifier
	Number
	StringLiteral
	Assign
	Plus
	Minus
	Star
	Slash
	Percent
	Amper
	Pipe
	Less
	Greater
	Bang
	Equal
	NotEqual
	Unknown
)

var kindOf = map[string]PieceType{
	// Keywords
	"yen":   Keyword,
	"sol":   Keyword,
	"sollu": Keyword,

	// Arithmetic operators
	"+": Plus,
	"-": Minus,
	"*": Star,
	"/": Slash,
	"%": Percent,

	// Logical operators
	"=": Assign,
	"<": Less,
	">": Greater,
	"!": Bang,
	"&": Amper,
	"|": Pipe,
	"==": Equal,
	"!=": Equal,

	"\n": Eol,
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
	case Assign:
		return fmt.Sprintf("assignment: %s", p.Value)
	case Plus:
		return fmt.Sprintf("plus: %s", p.Value)
	case Minus:
		return fmt.Sprintf("minus: %s", p.Value)
	case Star:
		return fmt.Sprintf("star: %s", p.Value)
	case Slash:
		return fmt.Sprintf("slash: %s", p.Value)
	case Percent:
		return fmt.Sprintf("percent: %s", p.Value)
	case Amper:
		return fmt.Sprintf("amper: %s", p.Value)
	case Pipe:
		return fmt.Sprintf("pipe: %s", p.Value)
	case Less:
		return fmt.Sprintf("less: %s", p.Value)
	case Greater:
		return fmt.Sprintf("greater: %s", p.Value)
	case Bang:
		return fmt.Sprintf("bang: %s", p.Value)
	case Equal:
		return fmt.Sprintf("assignment: %s", p.Value)
	case Eol:
		return "nextLine"
	case Eof:
		return "END"
	default:
		return fmt.Sprintf("unknown: %s", p.Value)
	}
}
