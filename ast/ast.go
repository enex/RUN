package ast

import (
	"fmt"
	"github.com/enex/RUN/token"
	"strconv"
)

//node is the interface vor virtualy everything
type Node interface {
	Pos() token.Position
	//String() string
}

type Paren struct {
	token.Position
	Content []Node
}

func (p Paren) String() string {
	r := "( "
	for _, e := range p.Content {
		r += fmt.Sprint(e, " ")
	}
	return r + ")"
}

type Number struct {
	token.Position
	Num float64
}

func (n Number) String() string {
	return fmt.Sprint(n.Num)
}

func NewNumber(s string, pos token.Position) Number {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return Number{
		Num:      num,
		Position: pos,
	}
}

type String struct {
	string
	token.Position
}

func (s String) String() string {
	//TODO: im String escapen
	return "\"" + s.string + "\""
}
func (s String) Value() string {
	return s.string
}

func NewString(s string, pos token.Position) String {
	return String{s, pos}
}

//Symbol welches als plazhalter benutzt wird
type Symbol struct {
	string
	token.Position
}

type Pattern interface{}

//Definition gleicht dem "="
type Definition struct {
	Pattern Pattern
	value   Node
	token.Position
}

//Ein Match mit einem Pattern
type Match struct {
	*Definition
	Values []Node
}

//Der scope ist eine schein struktur die alle wichtigen
//Informationen in sich trägt
type Scope struct {
	parent *Scope
	Defs   []Definition
}

//parent of the scope, if not defined it will be nil
func (s *Scope) Parent() *Scope {
	return s.parent
}
