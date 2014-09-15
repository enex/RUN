package main

import (
	"fmt"
	"strconv"
	"unicode"
)

/*
Die sprace setzt nur auf ein Prinzip auf, Pattern matching.

 - Kalmmern und Einrückungen werden als gruppierungen gesehen
 - Funktionen werden als Pattern definiert. Dieße sind ähnlich regulären
   Ausdrücken und ermöglichen das Defieren eines eigenen syntay. Die reihenfolge
   der Definitionen reflektiert die reihenfolge der match-vorgänge.

[Pattern] = [body]
*/

//Position der elemente
type Position struct {
	Line uint16 //Zeile im Dokument
	Col  uint8  //Zeichen in der Zeile
}

/*
func (p Position) String() string {
	return fmt.Sprint(p.Line, ":", p.Col)
}*/

func (p Position) Pos() Position {
	return p
}

type String struct {
	string
	Position
}

//Struktur die eine Quellcode-datei representiert
type File struct {
	path  string
	Scope //Der scope den dieße Datei besitzt
}

// Scanner
type Scanner struct {
	ch      rune //last charracter scanned
	offset  int  //offset of currnt character
	roffset int  //offset of next character

	pos Position

	src string //actual source code
}

// Init initializes Scanner and makes the source code ready to Scan
func (s *Scanner) Init(src string) {
	s.offset, s.roffset = 0, 0
	s.src = src

	s.next()
}

type Paren struct {
	Position
	content []Node
}

func (s *Scanner) scanParen() Node {
	sp := s.pos //safe the position
	s.next()    //go to the next char
	content := make([]Node, 0)
	for s.ch != ')' {
		fmt.Println("consume")
		content = append(content, s.scan())
	}
	s.next()
	return Paren{Position: sp, content: content}
}

func (s *Scanner) scan() Node {
	s.skipWhitespace()
	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}

	fmt.Println("scan", string(s.ch))

	switch s.ch {
	case '(':
		return s.scanParen()
	case ')':
		fmt.Println("RPAREN")
		panic("a paren has been closed without an opening paren")
	case '=':
		fmt.Println("definition")
	case '"': //String
		return s.scanString()
	default:
		if s.offset >= len(s.src)-1 {
			fmt.Println("ende erreicht")
		}

	}
	s.next()
	return nil
}

func (s *Scanner) Scan() Node {
	scope := make([]Node, 0)
	for !(s.offset >= len(s.src)-1) {
		n := s.scan()
		if n != nil {
			scope = append(scope, n)
		}
	}
	fmt.Println(scope)
	if len(scope) > 0 {
		return scope[0]
	}
	return nil
}

func (s *Scanner) next() {
	s.ch = rune(0)
	if s.roffset < len(s.src) {
		s.offset = s.roffset
		s.ch = rune(s.src[s.offset])
		if s.ch == '\n' {
			s.pos.Line++
			s.pos.Col = 0
		} else {
			s.pos.Col++
		}
		s.roffset++
	}
}

type Number struct {
	Position
	num float64
}

func NewNumber(s string, pos Position) Number {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return Number{
		num:      num,
		Position: pos,
	}
}

//Cosumes a number
func (s *Scanner) scanNumber() Node {
	start := s.offset
	sp := s.pos
	for unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	fmt.Println("scan Number")
	return NewNumber(s.src[start:offset], sp)
}

//Consumes one string
func (s *Scanner) scanString() Node {
	sp := s.pos //safe the position
	s.next()
	start := s.offset //start of the string
	//TODO: implement strings correctly
	for s.ch != '"' {
		if s.ch == rune(0) {
			break
		}
		s.next()
		if s.ch == '\\' { //skipp escape sequences
			s.next()
		}
	}
	offset := s.offset //end of the string
	s.next()
	//fmt.Println("String", s.src[start:offset], sp)
	return String{s.src[start:offset], sp}
}

/*
//scanns a comment, comments can actually be used by the compiler, so it
//will be scanned and saved but maybe it will be thrown away later in the
//Programm because no one cares about it
func (s *Scanner) scanComment() (string, Pos) {
	start := s.offset //start tof the comment
	for s.ch != '\n' && s.offset < len(s.src)-1 {
		s.next()
	}
	s.next()
	offset := s.offset + 1 //end of the comment
	return s.src[start:offset]
}*/

//will skipp all the white spaces without caring about them
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' {
		s.next()
	}
}

func (f *File) Read(s string) {
	sc := Scanner{}
	sc.Init(s)
	fmt.Println("Scanned: ", sc.Scan())
}

//Symbol welches als plazhalter benutzt wird
type Symbol struct {
	string
	Position
}

type Pattern interface{}

//node is the interface vor virtualy everything
type Node interface {
	Pos() Position
	//String() string
}

//Definition gleicht dem "="
type Definition struct {
	Pattern Pattern
	value   Node
	Position
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

func main() {
	f := &File{}
	f.Read(`
test = 
	Haus = "test"
	Auto = 
		if Haus == "test"
			"ja"
			"nein"
	Baum = "nicht da"
`)
}
