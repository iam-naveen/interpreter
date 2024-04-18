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

func initial(lex *Lexer) consumer {
	for {
		if lex.cur >= len(lex.input) {
			lex.send(Eof)
			return nil
		}
		if lex.takeOne("\n") {
			lex.send(Eol)
			continue
		}
		if lex.takeOne(" \t") {
			lex.ignore()
			continue
		}
		if lex.takeOne("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_") {
			return consumeAlphaNumeric
		}
		if lex.takeOne("0123456789") {
			return consumeNumber
		}
		if lex.takeOne("+*-/%") {
			val := string(lex.input[lex.start:lex.cur])
			lex.send(kindOf[val])
			continue
		}
		if lex.takeOne("=") {
			lex.send(Equal)
			continue
		} else {
			lex.next()
			lex.send(Unknown)
		}
	}
}
