package object

import (
	"bytes"
	"fmt"
	"strings"

	"monkey/ast"
)

type ObjectType string

const (
	INT_OBJ    = "INTEGER"
	BOOL_OBJ   = "BOOLEAN"
	NULL_OBJ   = "NULL"
	RETURN_OBJ = "RETURN_VALUE"
	ERROR_OBJ  = "ERROR"
    FUNCTION_OBJ = "FUNCTION"
    STRING_OBJ = "STRING"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return INT_OBJ }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOL_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

type Function struct {
    Parameters []*ast.Identifier
    Body *ast.BlockStatement
    Env *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
    var out bytes.Buffer
    
    params := []string{}
    for _, param := range f.Parameters {
        params = append(params, param.String())
    }

    out.WriteString("fn")
    out.WriteString("(")
    out.WriteString(strings.Join(params, ", "))
    out.WriteString(") {\n")    
    out.WriteString(f.Body.String())
    out.WriteString("\n}")
    
    return out.String()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "Error: " + e.Message }

type String struct {
    Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }
