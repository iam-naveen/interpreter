package lexer

import "fmt"

type PieceType int

const (
	Eof PieceType = iota
	Eol

	DataType
	If
	Else
	While
	For

	Print
	Input
	Length

	Identifier
	Number
	Boolean
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
	LessEqual
	GreaterEqual
	And
	Or

	ParanOpen
	ParanClose
	BraceOpen
	BraceClose
	BracketOpen
	BracketClose

	Unknown
)

var kindOf = map[string]PieceType{
	// Keywords
	"yen":       DataType,
	"sol":       DataType,
	"aam":       Boolean,
	"illai":     Boolean,
	"endral":    If,
	"illana":    Else,
	"varaikkum": While,
	"murai":     For,

	// builtins
	"sollu":     Print,
	"kodu":      Input,
	"neelam":    Length,

	// operators
	"+": Plus,
	"-": Minus,
	"*": Star,
	"/": Slash,
	"%": Percent,
	"=": Assign,

	// logical
	"<":  Less,
	">":  Greater,
	"!":  Bang,
	"&":  Amper,
	"|":  Pipe,
	"==": Equal,
	"!=": NotEqual,
	"<=": LessEqual,
	">=": GreaterEqual,

	// Grouping
	"(": ParanOpen,
	")": ParanClose,
	"{": BraceOpen,
	"}": BraceClose,
	"[": BracketOpen,
	"]": BracketClose,

	";": Eol,
}

type Piece struct {
	Kind  PieceType
	Value string
}

func (p Piece) String() string {
	switch p.Kind {
	case DataType:
		return fmt.Sprintf("keyword: %s", p.Value)
	case Identifier:
		return fmt.Sprintf("identifier: %s", p.Value)
	case Number:
		return fmt.Sprintf("number: %s", p.Value)
	case StringLiteral:
		return fmt.Sprintf("string: %s", p.Value)
	case Boolean:
		return fmt.Sprintf("boolean: %s", p.Value)
	case While:
		return fmt.Sprintf("while: %s", p.Value)
	case For:
		return fmt.Sprintf("for: %s", p.Value)
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
		return fmt.Sprintf("equal: %s", p.Value)
	case NotEqual:
		return fmt.Sprintf("not equal: %s", p.Value)
	case LessEqual:
		return fmt.Sprintf("less equal: %s", p.Value)
	case GreaterEqual:
		return fmt.Sprintf("greater equal: %s", p.Value)
	case And:
		return fmt.Sprintf("and: %s", p.Value)
	case Or:
		return fmt.Sprintf("or: %s", p.Value)
	case ParanOpen:
		return fmt.Sprintf("paran open: %s", p.Value)
	case ParanClose:
		return fmt.Sprintf("paran close: %s", p.Value)
	case BraceOpen:
		return fmt.Sprintf("brace open: %s", p.Value)
	case BraceClose:
		return fmt.Sprintf("brace close: %s", p.Value)
	case BracketOpen:
		return fmt.Sprintf("bracket open: %s", p.Value)
	case BracketClose:
		return fmt.Sprintf("bracket close: %s", p.Value)
	case Print:
		return fmt.Sprintf("print: %s", p.Value)
	case Input:
		return fmt.Sprintf("input: %s", p.Value)
	case If:
		return fmt.Sprintf("if: %s", p.Value)
	case Else:
		return fmt.Sprintf("else: %s", p.Value)
	case Eol:
		return ";"
	case Eof:
		return "END"
	default:
		return fmt.Sprintf("unknown: %s", p.Value)
	}
}
