package evaluator

import "fmt"


var builtins = map[string]func(args ...interface{}) interface{}{
	// print function
	"sollu": func(args ...interface{}) interface{} {
		for _, arg := range args {
			fmt.Println(arg)
		}
		return nil
	},
}
