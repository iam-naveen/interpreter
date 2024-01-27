package token

import "fmt"

type Type string

type Token struct {
	Type  Type
	Value string
}

const (
	// Keywords
	YEN       = "YEN"
	SOL       = "SOL"
	AAGA      = "AAGA"
	IRUNDHAAL = "IRUNDHAAL"
	ILLANA    = "ILLANA"
	IRUKUM    = "IRUKUM"
	ILLADHA   = "ILLADHA"
	MURAI     = "MURAI"
	VARAI     = "VARAI"
	SEIYAL    = "SEIYAL"

	// Literals
	ID  = "ID"
	INT = "INT"
	STR = "STR"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"
	STAR   = "*"
	SLASH  = "/"
	MOD    = "%"
	EQ     = "=="
	NEQ    = "!="
	LT     = "<"
	GT     = ">"
	LE     = "<="
	GE     = ">="
	AND    = "&&"
	OR     = "||"
	NOT    = "!"

	// Delimiters
	PIPE	  = "|"
	ARRROW    = "->"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	// Special
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	EOL     = "EOL"
)

var keywords = map[string]Type{
	"yen":      YEN,
	"sol":      SOL,
	"aaga":     AAGA,
	"irundhaal": IRUNDHAAL,
	"illana":   ILLANA,
	"irukum":   IRUKUM,
	"illadha":  ILLADHA,
	"murai":    MURAI,
	"varai":    VARAI,
	"seiyal":   SEIYAL,
}

// this function check if the given string
// is a keyword or not and returns the token type
func LookUp(ident string) Type {
	if keywordType, ok := keywords[ident]; ok {
		return keywordType
	}
	return ID
}

func (tok* Token) String() string {
	return fmt.Sprintf("Token(%s, %s)", tok.Type, tok.Value)
}
