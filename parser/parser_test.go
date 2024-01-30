package parser

import (
    "fmt"

	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 1000;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned 'nil'")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements not 3. got %d", len(program.Statements))
	}

	for i, tt := range []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	} {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
return 0;
return 10;
return 1235135;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements not 3. got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got %T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements not 1. got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got %T",
			program.Statements[0])
	}
    
    if !testIdentifier(t, stmt.Expression, "foobar") {
        return
    }
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "10;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements not 1. got %d", len(program.Statements))
	}
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got %T",
			program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("exp not *ast.IntegerLiteral. got %T", stmt.Expression)
	}
	if literal.Value != 10 {
		t.Errorf("literal.Value not %d. got %T", 10, literal.Value)
	}
	if literal.TokenLiteral() != "10" {
		t.Errorf("literal.TokenLiteral not %s. got %s", "10",
			literal.TokenLiteral())
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
    for _, tt := range []struct {
        input string
        operator string
        integerValue int64
    }{
        {"!5;", "!", 5},
        {"-15;", "-", 15},
    } {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 1 {
            t.Fatalf("program.Statements not 1. got %d", len(program.Statements))
        }

        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statements[0] not ast.ExpressionStatement. got %T",
                program.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.PrefixExpression)
        if !ok {
            t.Fatalf("exp.Operator not '%s'. got '%s'", tt.operator, exp.Operator)
        }
        if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
            return
        }
    }
}

func TestParsingInfixExpressions(t *testing.T) {
    for _, tt := range []struct {
        input string
        leftValue int64
        operator string
        rightValue int64
    }{
        {"5 + 5;", 5, "+", 5},
        {"5 - 5;", 5, "-", 5},
        {"5 * 5;", 5, "*", 5},
        {"5 / 5;", 5, "/", 5},
        {"5 > 5;", 5, ">", 5},
        {"5 < 5;", 5, "<", 5},
        {"5 == 5;", 5, "==", 5},
        {"5 != 5;", 5, "!=", 5},
    } {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        if len(program.Statements) != 1 {
            t.Fatalf("program.Statements not 1. got %d", len(program.Statements))
        }

        stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statements[0] not ast.ExpressionStatement. got %T",
                program.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.InfixExpression)
        if !ok {
            t.Fatalf("exp no ast.InfixExpression. got %T", stmt.Expression)
        }
        if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
            return
            }
        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator not '%s'. got %s", tt.operator, exp.Operator)
        }
        if !testIntegerLiteral(t, exp.Right, tt.rightValue) {
            return
        }
    }
}

func TestOperatorPrecedenceParsing(t *testing.T) {
    tests := []struct {
        input string
        expected string
    }{
        {
            "-a * b",
            "((-a) * b)",
        },
        {
            "!-a",
            "(!(-a))",
        },
        {
            "a + b + c",
            "((a + b) + c)",
        },
        {
            "a + b - c",
            "((a + b) - c)",
        },
        {
            "a * b * c",
            "((a * b) * c)",
        },
        {
            "a * b / c",
            "((a * b) / c)",
        },
        {
            "a + b / c",
            "(a + (b / c))",
        },
        {
            "a + b * c + d / e - f",
            "(((a + (b * c)) + (d / e)) - f)",
        },
        {
            "3 + 4; -5 * 5",
            "(3 + 4)((-5) * 5)",
        },
        {
            "5 > 4 == 3 < 4",
            "((5 > 4) == (3 < 4))",
        },
        {
            "5 < 4 != 3 > 4",
            "((5 < 4) != (3 > 4))",
        },
        {
            "3 + 4 * 5 == 3 * 1 + 4 * 5",
            "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
        },
        {
            "3 + 4 * 5 == 3 * 1 + 4 * 5",
            "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
        },
    }

    for _, tt := range tests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParserErrors(t, p)

        actual := program.String()

        if actual != tt.expected {
            t.Errorf("expected %q, got %q", tt.expected, actual)    
        }
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("%d parsing errors", len(errors))
	for _, msg := range errors {
		t.Errorf("error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got %q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got %T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got %s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got %s", name, letStmt.Name)
		return false
	}
	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
    integ, ok := il.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("il not *ast.IntegerLiteral. got %T", il)
        return false
    }

    if integ.Value != value {
        t.Errorf("integ.Value not %d. got %d", value, integ.Value)
        return false
    }

    if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
        t.Errorf("integ.TokenLiteral not %d. got %s", value,
            integ.TokenLiteral())
        return false
    }
    return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
   	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *ast.Identifier. got %T", exp)
        return false
	}
	if ident.Value != value {
		t.Errorf("ident.Value not %s. got %s", "foobar", ident.Value)
        return false
	}
	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got %s", "foobar", ident.TokenLiteral())
        return false
	}
    return true
}

func testLiteralExpression(
    t *testing.T,
    exp ast.Expression,
    expected interface{},
) bool {

    switch v := expected.(type) {
    case int:
        return testIntegerLiteral(t, exp, int64(v))
    case int64:
        return testIntegerLiteral(t, exp, v)
    case string:
        return testIdentifier(t, exp, v)
    }
    t.Errorf("type of exp not handled. got %T", exp)
    
    return false;
}

func testInfixExpression(
    t *testing.T,
    exp ast.Expression,
    left interface{},
    operator string,
    right interface{}) bool {

    opExp, ok := exp.(*ast.InfixExpression)
    if !ok {
        t.Errorf("exp is not ast.OperatorExpression. got %T(%s)", exp, exp)
        return false
    }
    if !testLiteralExpression(t, opExp.Left, left) {
        return false
    }
    if opExp.Operator != operator {
        t.Errorf("exp.Operator not '%s'. got %q", operator, opExp.Operator)
        return false
    }
    if !testLiteralExpression(t, opExp.Right, right) {
        return false
    }
    return true
}
