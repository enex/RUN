package ast_test

import (
	"testing"

	"github.com/enex/RUN/ast"
	"github.com/enex/RUN/token"
)

func TestBasicLit(t *testing.T) {
	b := &ast.BasicLit{
		LitPos: token.Pos(1),
		Kind:   token.INTEGER,
		Lit:    "24",
	}
	pos, end := token.Pos(1), token.Pos(3)
	if b.Pos() != pos || b.End() != end {
		t.Fatal("Expected:", pos, end, "Got:", b.Pos(), b.End())
	}
}

func TestBinaryExpr(t *testing.T) {
	// (+ 3 5)
	x := &ast.BasicLit{
		LitPos: token.Pos(4),
		Kind:   token.INTEGER,
		Lit:    "3",
	}
	y := &ast.BasicLit{
		LitPos: token.Pos(6),
		Kind:   token.INTEGER,
		Lit:    "5",
	}
	b := &ast.BinaryExpr{
		Expression: ast.Expression{
			Opening: token.Pos(1),
			Closing: token.Pos(7),
		},
		Op:    token.ADD,
		OpPos: token.Pos(2),
		List:  []ast.Expr{x, y},
	}

	if b.Pos() != token.Pos(1) {
		t.Fatal("BinaryExpr: Expected: 1 Got:", b.Pos())
	}
	if b.End() != token.Pos(7) {
		t.Fatal("BinaryExpr: Expected: 7 Got:", b.End())
	}
}
