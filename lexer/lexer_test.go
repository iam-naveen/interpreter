package lexer

import (
	"os"
	"testing"

	"github.com/iam-naveen/compiler/token"
)

func TestLexer(t *testing.T) {

	data, err := os.ReadFile("test_code")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		expectedType  token.Type
		expectedValue string
	}{
		{token.YEN, "yen"},
		{token.ID, "a"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.PLUS, "+"},
		{token.ID, "b"},
		{token.EOL, ""},
		{token.YEN, "yen"},
		{token.ID, "b"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
		{token.PLUS, "+"},
		{token.INT, "30"},
		{token.EOL, ""},
		{token.YEN, "yen"},
		{token.ID, "c"},
		{token.ASSIGN, "="},
		{token.ID, "a"},
		{token.PLUS, "+"},
		{token.ID, "b"},
		{token.EOL, ""},
		{token.SOL, "sol"},
		{token.ID, "name"},
		{token.ASSIGN, "="},
		{token.STR, "naveen"},
		{token.EOL, ""},
		{token.SOL, "sol"},
		{token.ID, "age"},
		{token.ASSIGN, "="},
		{token.ID, "b"},
		{token.EOL, ""},
		{token.SOL, "sol"},
		{token.ID, "new"},
		{token.ASSIGN, "="},
		{token.STR, ""},
		{token.EOL, ""},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.SLASH, "/"},
		{token.STAR, "*"},
		{token.MINUS, "-"},
		{token.EOL, ""},
		{token.ID, "a"},
		{token.LT, "<"},
		{token.ID, "b"},
		{token.AAGA, "aaga"},
		{token.IRUNDHAAL, "irundhaal"},
		{token.LBRACE, "{"},
		{token.EOL, ""},
		{token.ID, "a"},
		{token.ID, "sollu"},
		{token.EOL, ""},
		{token.RBRACE, "}"},
		{token.ILLANA, "illana"},
		{token.LBRACE, "{"},
		{token.EOL, ""},
		{token.ID, "b"},
		{token.ID, "sollu"},
		{token.EOL, ""},
		{token.RBRACE, "}"},
		{token.EOL, ""},
		{token.EOF, ""},
	}

	l := New(data)

	for i, test := range tests {
		token := l.NextToken()

		if token.Type != test.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, test, token)
		}

		if token.Value != test.expectedValue {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, test.expectedValue, token.Value)
		}
	}
}
