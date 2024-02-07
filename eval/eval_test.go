package eval

import (
	"testing"

	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

func TestEvalIntegerExpresion(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"20 + 2 * -10", 0},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * (3 * 3) + 10", 37},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	} {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestEvalBooleanExpression(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	} {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestBangOperator(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	} {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpressions(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 < 2) { 10 } else { 20 }", 10},
	} {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestReturnStatements(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
	} {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operation: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operation: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operation: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"undefined identifier: foobar",
		},
	} {
		evaluated := testEval(tt.input)

		errObj, ok := evaluated.(*object.Error)
		if !ok {
			t.Errorf("Object not Error. Got %T (%+v)",
				evaluated, evaluated)
			continue
		}
		if errObj.Message != tt.expected {
			t.Errorf("Message mismatch. Expected %q, got %q",
				tt.expected, errObj.Message)
		}
	}
}

func TestLetStatement(t *testing.T) {
	for _, tt := range []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	} {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
    input := "fn(x) { x + 2; };"

    evaluated := testEval(input)
    fn, ok := evaluated.(*object.Function)
    if !ok {
        t.Fatalf("Object not Function. got %T (%+v)", evaluated, evaluated)
    }

    if len(fn.Parameters) != 1 {
        t.Fatalf("Wrong amount of parameters. Parameters: %+v",
            fn.Parameters)
    }
    if fn.Parameters[0].String() != "x" {
        t.Fatalf("Parameter not 'x'. Got %q", fn.Parameters[0])
    }

    expectedBody := "(x + 2)"
    if fn.Body.String() != expectedBody {
        t.Fatalf("Body mismatch. Expected %q, got %q",
            expectedBody, fn.Body.String())
    }
}

func TestFunctionApplication(t *testing.T) {
    for _, tt := range []struct {
        input string
        expected int64
    }{
        {"let identity = fn(x) { x; }; identity(5);", 5},
        {"let identity = fn(x) { return x; }; identity(5);", 5},
        {"let double = fn(x) { x * 2; }; double(5);", 10},
        {"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
        {"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
        {"fn(x) { x; }(5)", 5},
    } {
        testIntegerObject(t, testEval(tt.input), tt.expected)
    }
}

func TestClosures(t *testing.T) {
    input := `
let newAdder = fn(x) {
    fn (y) { x + y };
};

let addTwo = newAdder(2);
addTwo(2);
`
    testIntegerObject(t, testEval(input), 4)
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnv()

	return Eval(program, env)
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

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("Object not Null. Got %T (%+v)", obj, obj)
		return false
	}
	return true
}
