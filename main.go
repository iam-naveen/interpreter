package main

import (
	"fmt"
	"os"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the input file")
		return
	}
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading the file")
		return
	}
	_, channel := lexer.CreateLexer(input)
	parser.Run(channel)
}
