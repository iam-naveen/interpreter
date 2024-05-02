package evaluator

import (
	"fmt"

	"github.com/iam-naveen/compiler/object"
)


var Builtins = map[string]func(args ...interface{}) interface{}{
	// print function
	"sollu": func(args ...interface{}) interface{} {
		for _, arg := range args {
			fmt.Println(arg)
		}
		return nil
	},
	// input function
	"kodu": func(args ...interface{}) interface{} {
		if len(args) > 1 {
			fmt.Println("kodu function takes only one argument")
			return nil
		}
		if args[0] != nil {
			switch t := args[0].(type) {
			case *object.String:
				fmt.Scanln(&t.Value)
				return t
			case *object.Integer:
				fmt.Scanln(&t.Value)
				return t
			default:
				fmt.Println("kodu function takes only one argument of type INTEGER or STRING")
				return nil
			}
		} 	
		return nil
	},
}
