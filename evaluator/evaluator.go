package evaluator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/iam-naveen/compiler/lexer"
	"github.com/iam-naveen/compiler/object"
	"github.com/iam-naveen/compiler/tree"
)

var console = bufio.NewReader(os.Stdin)

func Eval(node tree.Node, env *object.Environment) {
	switch node := node.(type) {

	// Statements
	case *tree.Program:
		evalProgram(node, env)
	case *tree.Block:
		evalBlock(node, env)
	case *tree.PrintStmt:
		evalPrintStmt(node, env)
	case *tree.Input:
		evalInput(node, env)
	case *tree.Declaration:
		evalDeclaration(node, env)
	case *tree.IfStmt:
		evalIfStatement(node, env)
	case *tree.WhileStmt:
		evalWhileStatement(node, env)
	case *tree.ForStmt:
		evalForStatement(node, env)
	case *tree.ExpressionStmt:
		evalExpressionStatement(node, env)

	default:
		err := &object.Error{Message: fmt.Sprintf("Unknown Node %T", node)}
		fmt.Println(err.Inspect())

	}
}

func evalInput(input *tree.Input, env *object.Environment) {

	switch input.DataType {
	case "INTEGER":
		evalIntegerInput(input, env)
	case "STRING":
		evalStringInput(input, env)
	default:
		fmt.Println("ERROR: Invalid Input")
		os.Exit(1)
	}
}

func evalIntegerInput(input *tree.Input, env *object.Environment) {
	var value int64
	fmt.Print(input.Variable.Name, " = ")
	_, err := fmt.Scanf("%d", &value)
	if err != nil {
		fmt.Println("ERROR: Invalid Input")
		os.Exit(1)
	}
	env.Set(input.Variable.Name, &object.Integer{Value: value})
}

func evalStringInput(input *tree.Input, env *object.Environment) {
	fmt.Print(input.Variable.Name, " = ")
	str, err := console.ReadString('\n')
	str = str[:len(str)-1] // Remove the delimeter '\n'
	if err != nil {
		fmt.Println("ERROR: Invalid Input")
		os.Exit(1)
	}
	env.Set(input.Variable.Name, &object.String{Value: str})
}

func evalProgram(program *tree.Program, env *object.Environment) {
	for _, stmt := range program.Children {
		Eval(stmt, env)
	}
}

func evalBlock(block *tree.Block, env *object.Environment) {
	for _, stmt := range block.Children {
		Eval(stmt, env)
	}
}

func evalWhileStatement(stmt *tree.WhileStmt, env *object.Environment) {
	for evaluateExpression(stmt.Condition, env).(*object.Boolean).Value {
		Eval(stmt.Body, env)
	}
}

func evalForStatement(stmt *tree.ForStmt, env *object.Environment) {
	count := evaluateExpression(stmt.Count, env)
	switch count := count.(type) {
	case *object.Integer:
		for i := int64(0); i < count.Value; i++ {
			Eval(stmt.Body, env)
		}
	default:
		fmt.Println("ERROR: Expected Constant Expression in For loop")
		os.Exit(1)
	}

}

func evalPrintStmt(stmt *tree.PrintStmt, env *object.Environment) {
	result := evaluateExpression(stmt.Value, env)
	fmt.Println(result.Inspect())
}

func evalDeclaration(decl *tree.Declaration, env *object.Environment) {
	value := evaluateExpression(decl.Value, env)
	if value.Type() == object.ERROR_OBJ {
		fmt.Println(value.Inspect())
		os.Exit(1)
	}
	if decl.Datatype == string(value.Type()) {
		env.Set(decl.Name.Value, value)
	} else {
		fmt.Println("ERROR: Cannot Assign", value.Type(), "to", decl.Datatype, "variable")
		os.Exit(1)
	}
}

func evalExpressionStatement(stmt *tree.ExpressionStmt, env *object.Environment) {
	switch expr := stmt.Expression.(type) {
	case *tree.Assign:
		evalAssign(expr, env)
	}
}

func evalIfStatement(stmt *tree.IfStmt, env *object.Environment) {
	result := evaluateExpression(stmt.Condition, env)
	if result.Type() != object.BOOLEAN_OBJ {
		err := &object.Error{Message: "Non Boolean Expression in If Statement"}
		fmt.Println(err.Inspect())
		return
	}
	if result.(*object.Boolean).Value {
		Eval(stmt.Then, env)
	} else if stmt.Else != nil {
		Eval(stmt.Else, env)
	}
}

func evalAssign(assign *tree.Assign, env *object.Environment) {
	value := evaluateExpression(assign.Right, env)
	env.Set(assign.Left.Name, value)
}

func evaluateExpression(expr tree.Expr, env *object.Environment) object.Object {
	switch expr := expr.(type) {
	case *tree.Number:
		return &object.Integer{Value: expr.Value}
	case *tree.StringLiteral:
		return &object.String{Value: expr.Value}
	case *tree.Boolean:
		return &object.Boolean{Value: expr.Value}
	case *tree.Identifier:
		res, ok := env.Get(expr.Name)
		if !ok {
			return &object.Error{Message: "Unknown identifier"}
		}
		return res
	case *tree.Access:
		left := evaluateExpression(expr.Left, env)
		index := evaluateExpression(expr.Index, env)
		switch left := left.(type) {
		case *object.String:
			if index.Type() != object.INTEGER_OBJ {
				return &object.Error{Message: "Index must be an Integer"}
			}
			i := index.(*object.Integer).Value
			if i < 0 || i >= int64(len(left.Value)) {
				return &object.Error{Message: "Index out of range"}
			}
			return &object.String{Value: string(left.Value[i])}
		}
	case *tree.Length:
		value := evaluateExpression(expr.Value, env)
		switch value := value.(type) {
		case *object.String:
			return &object.Integer{Value: int64(len(value.Value))}
		default:
			return &object.Error{Message: "Length can only be applied to Strings"}
		}
	case *tree.Prefix:
		right := evaluateExpression(expr.Right, env)
		switch expr.Operator.Kind {
		case lexer.Minus:
			if right.Type() != object.INTEGER_OBJ {
				return &object.Error{Message: "Invalid Operand Type"}
			}
			return &object.Integer{Value: -right.(*object.Integer).Value}
		case lexer.Plus:
			if right.Type() != object.INTEGER_OBJ {
				return &object.Error{Message: "Invalid Operand Type"}
			}
			return right
		case lexer.Bang:
			if right.Type() != object.BOOLEAN_OBJ {
				return &object.Error{Message: "Invalid Operand Type"}
			}
			return &object.Boolean{Value: !right.(*object.Boolean).Value}
		}
	case *tree.Binary:
		left := evaluateExpression(expr.Left, env)
		right := evaluateExpression(expr.Right, env)
		switch expr.Operator.Kind {
		case lexer.Plus:
			return evaluatePlus(left, right)
		case lexer.Minus:
			return evaluateMinus(left, right)
		case lexer.Star:
			return evaluateMultiply(left, right)
		case lexer.Slash:
			return evaluateDivide(left, right)
		case lexer.Percent:
			return &object.Integer{Value: left.(*object.Integer).Value % right.(*object.Integer).Value}
		case lexer.Equal:
			return &object.Boolean{Value: left.Inspect() == right.Inspect()}
		case lexer.NotEqual:
			return &object.Boolean{Value: left.Inspect() != right.Inspect()}
		case lexer.Less:
			return &object.Boolean{Value: left.(*object.Integer).Value < right.(*object.Integer).Value}
		case lexer.Greater:
			return &object.Boolean{Value: left.(*object.Integer).Value > right.(*object.Integer).Value}
		case lexer.LessEqual:
			return &object.Boolean{Value: left.(*object.Integer).Value <= right.(*object.Integer).Value}
		case lexer.GreaterEqual:
			return &object.Boolean{Value: left.(*object.Integer).Value >= right.(*object.Integer).Value}
		case lexer.And:
			return &object.Boolean{Value: left.(*object.Boolean).Value && right.(*object.Boolean).Value}
		case lexer.Or:
			return &object.Boolean{Value: left.(*object.Boolean).Value || right.(*object.Boolean).Value}
		default:
			msg := fmt.Sprintf(
				"Unknown Operator '%s' for %s",
				expr.Operator.Value,
				left.Type(),
			)
			return &object.Error{Message: msg}
		}
	}
	return &object.Error{Message: "Unknown expression"}
}

func evaluatePlus(left, right object.Object) object.Object {
	switch left := left.(type) {
	case *object.Integer:
		if right.Type() == object.STRING_OBJ {
			return &object.String{Value: fmt.Sprintf("%d%s", left.Value, right.(*object.String).Value)}
		}
		return &object.Integer{Value: left.Value + right.(*object.Integer).Value}
	case *object.String:
		if right.Type() == object.INTEGER_OBJ {
			return &object.String{Value: left.Value + fmt.Sprintf("%d", right.(*object.Integer).Value)}
		}
		return &object.String{Value: string(left.Value) + right.(*object.String).Value}
	default:
		return &object.Error{Message: "Unknown Type"}
	}
}

func evaluateMinus(left, right object.Object) object.Object {
	if left.Type() != right.Type() {
		return &object.Error{Message: fmt.Sprintf("Type Mismatch: Cannot perform operation with %s and %s", left.Type(), right.Type())}
	}
	switch left := left.(type) {
	case *object.Integer:
		return &object.Integer{Value: left.Value - right.(*object.Integer).Value}
	case *object.String:
		return &object.Error{Message: "Cannot Subtract Strings"}
	default:
		return &object.Error{Message: "Unknown Type"}
	}
}

func evaluateMultiply(left, right object.Object) object.Object {
	if left.Type() != right.Type() {
		return &object.Error{Message: fmt.Sprintf("Type Mismatch: Cannot perform operation with %s and %s", left.Type(), right.Type())}
	}
	switch left := left.(type) {
	case *object.Integer:
		return &object.Integer{Value: left.Value * right.(*object.Integer).Value}
	case *object.String:
		return &object.Error{Message: "Cannot Multiply Strings"}
	default:
		return &object.Error{Message: "Unknown Type"}
	}
}

func evaluateDivide(left, right object.Object) object.Object {
	if left.Type() != right.Type() {
		return &object.Error{Message: fmt.Sprintf("Type Mismatch: Cannot perform operation with %s and %s", left.Type(), right.Type())}
	}
	switch left := left.(type) {
	case *object.Integer:
		return &object.Integer{Value: left.Value / right.(*object.Integer).Value}
	case *object.String:
		return &object.Error{Message: "Cannot Divide Strings"}
	default:
		return &object.Error{Message: "Unknown Type"}
	}
}
