package ast

import "reflect"

type Func func(Node)

func Walk(node Node, f Func) {
	if node == nil || reflect.ValueOf(node).IsNil() {
		return
	}

	if f != nil {
		f(node)
	}
	switch n := node.(type) {
	case *BinaryExpr:
		for _, v := range n.List {
			Walk(v, f)
		}
	case *CallExpr:
		Walk(n.Name, f)
		for _, v := range n.Args {
			Walk(v, f)
		}
	case *DeclExpr:
		Walk(n.Name, f)
		for _, v := range n.Params {
			Walk(v, f)
		}
		Walk(n.Type, f)
		Walk(n.Body, f)
	case *ExprList:
		for _, v := range n.List {
			Walk(v, f)
		}
	case *File:
		for _, v := range n.Scope.Table {
			Walk(v.Value, f)
		}
	case *Package:
		for _, v := range n.Scope.Table {
			Walk(v.Value, f)
		}
	case *UnaryExpr:
		Walk(n.Value, f)
	}
}
