package eval

import (
    "testing"

    "monkey/lexer"
    "monkey/object"
    "monkey/parser"
)

func TestEvalIntegerExpresion(t *testing.T) {
    for _, tt := range []struct {
        input string
        expected int64
    }{
        {"5", 5},
        {"10", 10},
    } {
        evaluated := testEval(tt.input)
        testIntegerObject(t, evaluated, tt.expected)
    }
}

func TestEvalBooleanExpression(t *testing.T) {
    for _, tt := range []struct {
        input string
        expected bool
    }{
        {"true", true},
        {"false", false},
    } {
        evaluated := testEval(tt.input)
        testBooleanObject(t, evaluated, tt.expected)
    }
}

func testEval(input string) object.Object {
    l := lexer.New(input)
    p := parser.New(l)
    program := p.ParseProgram()

    return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
    result, ok := obj.(*object.Integer)
    if !ok {
        t.Errorf("Object not Integer. Got %T (%+v).", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("Object.Value mismatch. Expected %d, got %d",
            expected, result.Value)
        return false
    }
    return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
    result, ok := obj.(*object.Boolean)
    if !ok {
        t.Errorf("Object not Boolean. Got %T (%+v)", obj, obj)
        return false
    }
    if result.Value != expected {
        t.Errorf("Object.Value mismatch. Expected %t, got %t",
            expected, result.Value)
        return false
    }
    return true
}
