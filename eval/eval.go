package eval

import (
    "fmt"

	"monkey/ast"
	"monkey/object"
)

var (
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
	NULL  = &object.Null{}
)

func Eval(node ast.Node) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements)
	case *ast.BlockStatement:
		return evalBlockStatement(node)
    case *ast.ReturnStatement:
        val := Eval(node.ReturnValue)
        if isError(val) {
            return val
        }
        return &object.ReturnValue{Value: val}
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return booleanObject(node.Value)
	case *ast.PrefixExpression:
		right := Eval(node.Right)
        if isError(right) {
            return right
        }
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left)
        if isError(left) {
            return left
        }
		right := Eval(node.Right)
        if isError(right) {
            return right
        }
		return evalInfixExpression(node.Operator, left, right)
	case *ast.IfExpression:
		return evalIfExpression(node)
	}

	return nil
}

func evalProgram(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
        result = Eval(stmt)

        switch result := result.(type) {
        case *object.ReturnValue:
            return result.Value
        case *object.Error:
            return result
        }
	}
	return result
}

func evalBlockStatement(block *ast.BlockStatement) object.Object {
    var result object.Object

    for _, stmt := range block.Statements {
        result = Eval(stmt)

        if result != nil {
            rt := result.Type()
            if rt == object.RETURN_OBJ || rt == object.ERROR_OBJ {
                return result
            }
        }
    }
    return result
}

func booleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangPrefixOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	default:
        return newError("unknown operation: %s%s", operator, right.Type())
	}
}

func evalBangPrefixOperator(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.INT_OBJ {
        return newError("unknown operation: -%s", right.Type())
	}
	value := right.(*object.Integer).Value

	return &object.Integer{Value: -value}
}

func evalInfixExpression(
	operator string,
	left object.Object,
	right object.Object,
) object.Object {

	switch {
	case left.Type() == object.INT_OBJ && right.Type() == object.INT_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "==":
		return booleanObject(left == right)
	case operator == "!=":
		return booleanObject(left != right)
    case left.Type() != right.Type():
        return newError("type mismatch: %s %s %s",
            left.Type(), operator, right.Type())
	default:
        return newError("unknown operation: %s %s %s",
            left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(
	operator string,
	left object.Object,
	right object.Object,
) object.Object {

	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return booleanObject(leftVal < rightVal)
	case ">":
		return booleanObject(leftVal > rightVal)
	case "==":
		return booleanObject(leftVal == rightVal)
	case "!=":
		return booleanObject(leftVal != rightVal)
	default:
        return newError("unknown operation: %s %s %s",
            left.Type(), operator, right.Type())
	}
}

func evalIfExpression(ie *ast.IfExpression) object.Object {
	condition := Eval(ie.Condition)
    if isError(condition) {
        return condition
    }

	if isTrue(condition) {
		return Eval(ie.Consequence)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative)
	}
	return NULL
}

func isTrue(obj object.Object) bool {
	switch obj {
	case FALSE, NULL:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
    return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
    return obj != nil && obj.Type() == object.ERROR_OBJ
}
