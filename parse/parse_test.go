package parse_test

import (
	"testing"

	"github.com/enex/RUN/ast"
	"github.com/enex/RUN/parse"
)

type Test struct {
	name  string
	src   string
	types []Type
	pass  bool
}

type Type int

const (
	BASIC Type = iota
	BINARY
	CALL
	DEF
	IDENT
	IF
	FILE
	LIST
	UNARY
	UNKNOWN
	VAR
)

var typeStrings = []string{
	BASIC:   "basiclit",
	BINARY:  "binaryexpr",
	CALL:    "callexpr",
	DEF:     "defexpr",
	IDENT:   "ident",
	IF:      "if",
	FILE:    "file",
	LIST:    "exprlist",
	UNARY:   "unaryexpr",
	UNKNOWN: "unknown",
	VAR:     "var",
}

func (t Type) String() string { return typeStrings[int(t)] }

func nodeTest(types []Type, t *testing.T) func(node ast.Node) {
	typ := UNKNOWN
	i := 0
	return func(node ast.Node) {
		switch node.(type) {
		case *ast.BasicLit:
			typ = BASIC
		case *ast.BinaryExpr:
			typ = BINARY
		case *ast.CallExpr:
			typ = CALL
		case *ast.DeclExpr:
			typ = DEF
		case *ast.ExprList:
			typ = LIST
		case *ast.File:
			typ = FILE
		case *ast.Ident:
			t.Log("ident:", node.(*ast.Ident).Name)
			typ = IDENT
		case *ast.UnaryExpr:
			typ = UNARY
		}
		if i > len(types)-1 {
			t.Fatal("index out of range")
		}
		if types[i] != typ {
			t.Fatal("Walk index:", i, "Expected:", types[i], "Got:", typ)
		}
		i++
	}
}

func handleTests(t *testing.T, tests []Test) {
	for _, test := range tests {
		f, _ := parse.ParseExpression(test.name, test.src)
		if f == nil && test.pass {
			t.Log(f == nil)
			t.Log(!test.pass)
			t.Fatal("Failed to parse")
		}
		ast.Walk(f, nodeTest(test.types, t))
	}
}

func TestParseBasic(t *testing.T) {
	tests := []Test{
		{"basic1", "24", []Type{BASIC}, true},
		{"basic2", "a", []Type{IDENT}, true},
	}
	handleTests(t, tests)
}

//Test the basic functionality
func TestParseBinary(t *testing.T) {
	tests := []Test{
		{"basic1", "(+ 2 3)", []Type{CALL, BASIC, BASIC}, true},
		{"basic2", "(+ 2 b)", []Type{CALL, BASIC, IDENT}, true},
		{"basic3", "(+ a b)", []Type{CALL, IDENT, IDENT}, true},
	}
	handleTests(t, tests)
}

func TestParseCall(t *testing.T) {
	tests := []Test{
		{"call1", "(add 1 2)", []Type{CALL, IDENT, BASIC, BASIC}, true},
		{"call2", "(nothing)", []Type{CALL, IDENT}, true},
	}
	handleTests(t, tests)
}

func TestParseComment(t *testing.T) {
	tests := []Test{
		{"comment1", "2; comment", []Type{BASIC}, true},
		{"comment2", "2; comment\na", []Type{BASIC, IDENT}, true},
		{"comment3", "; comment\na", []Type{IDENT}, true},
		{"comment4", "(+ 2; comment\n3)", []Type{BINARY, BASIC, BASIC}, true},
		{"comment5", ";comment", []Type{}, false},
	}
	handleTests(t, tests)
}

func TestParseDecl(t *testing.T) {
	tests := []Test{
		{"decl1", "(def func int 0)",
			[]Type{DEF, IDENT, IDENT, BASIC}, true},
		{"decl2", "(def five int (+ 2 3))",
			[]Type{DEF, IDENT, IDENT, BINARY, BASIC, BASIC}, true},
		{"decl3", "(def add(a b int) int (+ a b))",
			[]Type{DEF, IDENT, IDENT, IDENT, IDENT, BINARY, IDENT, IDENT}, true},
		{"decl4", "(def main () int a)", []Type{}, false},
		{"decl5", "(def main int ())", []Type{}, false},
		{"decl6", "def main int)", []Type{}, false},
	}
	handleTests(t, tests)
}
