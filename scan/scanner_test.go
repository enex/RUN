package scan

import (
	"fmt"
	"github.com/enex/RUN/ast"
	"testing"
)

func TestNumber(t *testing.T) {
	s := New("8117")
	r := s.scan()
	if r.(ast.Number).Num != 8117 {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	s := New(`"Hallo \" da"`)
	r := s.scan()
	t.Log(r)
	t.Log(r.(ast.String).Value())
	if r.(ast.String).Value() != "Hallo \" da" {
		t.Fail()
	}
}

func TestComment(t *testing.T) {
	s := New(`//Das ist ein Kommentar`)
	r := s.scan()
	t.Log(r)
	if r != nil {
		t.Fail()
	}
}

func TestParen(t *testing.T) {
	s := New("(23 34 75)")
	r := s.scan()
	t.Log(r)
	if fmt.Sprint(r) != "( 23 34 75 )" {
		t.Fail()
	}
}
/*
func TestIdent(t *testing.T) {
	s := New(`1
	21 22 23
		31 32 33
		321 322 323
	24 25
25
$a + $b = a
`)
	r := s.Scan()
	t.Log(r)
	t.Fail()
}*/
