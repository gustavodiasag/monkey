package eval

import (
    "monkey/ast"
    "monkey/object"
)

var (
    TRUE = &object.Boolean{Value: true}
    FALSE = &object.Boolean{Value: false}
    NULL = &object.Null{}
)

func Eval(node ast.Node) object.Object {
    switch node := node.(type) {
    case *ast.Program:
        return evalStatements(node.Statements)
    case *ast.ExpressionStatement:
        return Eval(node.Expression)
    case *ast.IntegerLiteral:
        return &object.Integer{Value: node.Value}
    case *ast.Boolean:
        return booleanObject(node.Value)
    case *ast.PrefixExpression:
        right := Eval(node.Right)
        return evalPrefixExpression(node.Operator, right)
    }
    
    return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
    var result object.Object

    for _, stmt := range stmts {
        result = Eval(stmt)
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
        return NULL
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
        return NULL
    }
    value := right.(*object.Integer).Value

    return &object.Integer{Value: -value}
}
