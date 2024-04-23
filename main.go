package main

import (
	"fmt"
	"os"
	"slices"

	"github.com/iam-naveen/compiler/evaluator"
	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/object"
	"github.com/iam-naveen/compiler/parser"
	"github.com/sanity-io/litter"
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
	ast := parser.Parse(channel, logging)
	if logging {
		litter.Dump(ast)
	}
	if slices.Contains(os.Args, "-t") || slices.Contains(os.Args, "--tree") {
		fmt.Println(ast.Print(0, "", ""))
	}
	env := object.NewEnvironment()
	evaluator.Eval(ast, env)
}
