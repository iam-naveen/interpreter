package lexer

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

func (lex *Lexer) run() {
	for consumer := initial; consumer != nil; {
		consumer = consumer(lex)
	}
	close(lex.channel)
}

const (
	alpha   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	numeric = "0123456789"
)

func initial(lex *Lexer) consumer {
	for {
		if lex.cur >= len(lex.input) {
			lex.send(Eof)
			return nil
		}
		if lex.takeOne(";") {
			lex.send(Eol)
			continue
		}
		if lex.takeOne(" \n\t") {
			lex.takeMany(" \n\t")
			lex.ignore()
			continue
		}
		if lex.takeOne(alpha) {
			return consumeAlphaNumeric
		}
		if lex.takeOne("0123456789") {
			lex.takeMany(numeric)
			lex.send(Number)
			continue
		}
		if lex.takeOne("\"") {
			return consumeString
		}
		if lex.takeOne("+*-/%(){}[]") {
			val := string(lex.input[lex.start:lex.cur])
			lex.send(kindOf[val])
			continue
		}
		if lex.takeOne("=!<>&|") {
			switch lex.input[lex.start] {
			case '=':
				lex.sendIfElse(lex.takeOne("="), Equal, Assign)
			case '!':
				lex.sendIfElse(lex.takeOne("="), NotEqual, Bang)
			case '<':
				lex.sendIfElse(lex.takeOne("="), LessEqual, Less)
			case '>':
				lex.sendIfElse(lex.takeOne("="), GreaterEqual, Greater)
			case '&':
				lex.sendIfElse(lex.takeOne("&"), And, Amper)
			case '|':
				lex.sendIfElse(lex.takeOne("|"), Or, Pipe)
			}
		} else {
			lex.next()
			lex.send(Unknown)
		}
	}
}
