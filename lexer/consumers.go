package lexer

type consumer func(*Lexer) consumer

func consumeAlphaNumeric(lex *Lexer) consumer {
	lex.takeMany(alpha + numeric)
	word := string(lex.input[lex.start:lex.cur])
	if kind, ok := kindOf[word]; ok {
		lex.send(kind)
	} else {
		lex.send(Identifier)
	}
	return initial
}

func consumeString(lex *Lexer) consumer {
	lex.ignore() // consume the opening "
	for lex.input[lex.cur] != '"' {
		lex.next()
		if lex.cur >= len(lex.input) {
			lex.send(Unknown)
			return nil
		}
	}
	lex.send(StringLiteral)

	// consume the closing "
	lex.next()
	lex.ignore()

	return initial
}
