package console

import (
	"bufio"
	"io"
	"fmt"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/token"
)

func Run(reader io.Reader, writer io.Writer) {
	input := bufio.NewScanner(reader)

	for {
		writer.Write( []byte("\n>> ") )
		running := input.Scan()
		if !running {
			return
		}
		line := input.Bytes()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOL; tok = l.NextToken() {
			if tok.Type == token.EOF {
				break
			}
			val := fmt.Sprintf("%+v\n", tok)
			writer.Write([]byte(val))
		}
	}
}
