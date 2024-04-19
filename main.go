package main

import (
	"fmt"
	"os"
	"slices"

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
	logging := false
	if len(os.Args) > 2 {
		if slices.Contains(os.Args, "-l") || slices.Contains(os.Args, "--log") {
			logging = true
		}
	}
	_, channel := lexer.CreateLexer(input)
	parser.Run(channel, logging)
}
