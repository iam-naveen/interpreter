package lexer

type consumer func(*Lexer) consumer

func consumeAlphaNumeric(lex *Lexer) consumer {
	lex.takeMany("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_")
	word := string(lex.input[lex.start:lex.cur])
	if kind, ok := kindOf[word]; ok {
		lex.send(kind)
	} else {
		lex.send(Identifier)
	}
	return initial
}

func consumeNumber(lex *Lexer) consumer {
	lex.takeMany("0123456789")
	lex.send(Number)
	return initial
}

