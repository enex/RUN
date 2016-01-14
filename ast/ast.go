package ast

import (
	"github.com/enex/RUN/token"
)

type Node interface {
	Pos() token.Pos
	End() token.Pos
}

type Expr interface {
	Node
	exprNode()
}

type BasicLit struct {
	LitPos token.Pos
	Kind   token.Token
	Lit    string
}

type BinaryExpr struct {
	Expression
	Op    token.Token
	OpPos token.Pos
	List  []Expr
}

type CallExpr struct {
	Expression
	Name *Ident
	Args []Expr
}

type DeclExpr struct {
	Expression
	Decl   token.Pos
	Name   *Ident
	Type   *Ident
	Params []*Ident
	Body   Expr
	Scope  *Scope
}

type Expression struct {
	Opening token.Pos
	Closing token.Pos
}

type ExprList struct {
	Expression
	List []Expr
}

type File struct {
	Scope *Scope
}

type Ident struct {
	NamePos token.Pos
	Name    string
	Object  *Object // may be nil (ie. Name is a type keyword)
}

type Object struct {
	NamePos token.Pos
	Name    string
	Kind    ObKind
	Offset  int
	Type    *Ident // variable type, function return type, etc
	Value   Expr
}

type ObKind int

type Package struct {
	Scope *Scope
	Files []*File
}

type Scope struct {
	Parent *Scope
	Table  map[string]*Object
}

type UnaryExpr struct {
	OpPos token.Pos
	Op    string
	Value Expr
}

func (b *BasicLit) Pos() token.Pos   { return b.LitPos }
func (e *Expression) Pos() token.Pos { return e.Opening }
func (f *File) Pos() token.Pos       { return token.NoPos }
func (i *Ident) Pos() token.Pos      { return i.NamePos }
func (p *Package) Pos() token.Pos    { return token.NoPos }
func (u *UnaryExpr) Pos() token.Pos  { return u.OpPos }

func (b *BasicLit) End() token.Pos   { return b.LitPos + token.Pos(len(b.Lit)) }
func (e *Expression) End() token.Pos { return e.Closing }
func (f *File) End() token.Pos       { return token.NoPos }
func (i *Ident) End() token.Pos      { return i.NamePos + token.Pos(len(i.Name)) }
func (p *Package) End() token.Pos    { return token.NoPos }
func (u *UnaryExpr) End() token.Pos  { return u.Value.End() }

func (b *BasicLit) exprNode()   {}
func (e *Expression) exprNode() {}
func (i *Ident) exprNode()      {}
func (u *UnaryExpr) exprNode()  {}

const (
	Decl ObKind = iota
	Var
)

func NewScope(parent *Scope) *Scope {
	return &Scope{Parent: parent, Table: make(map[string]*Object)}
}

func (s *Scope) Insert(ob *Object) *Object {
	if old, ok := s.Table[ob.Name]; ok {
		return old
	}
	s.Table[ob.Name] = ob
	return nil
}

func (s *Scope) Lookup(ident string) *Object {
	ob, ok := s.Table[ident]
	if ok || s.Parent == nil {
		return ob
	}
	return s.Parent.Lookup(ident)
}

func (s *Scope) Size() int {
	return len(s.Table)
}
