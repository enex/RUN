package main

import (
	"fmt"
	//"unicode"
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

func (s *Scanner) Scan() Node {
	s.skipWhitespace()
	lit := ""
	if s.ch == '\n' { //Parse line break and ident after linebreak
		lit = string(s.ch)
		s.next()
		for s.ch == '\t' {
			lit += string(s.ch)
			s.next()
		}
		return nil
	}

	/*
		if unicode.IsLetter(s.ch) {
			return s.scanIdentifier()
		}
	*/ /*
		if unicode.IsDigit(s.ch) {
			return s.scanNumber()
		}*/

	ch := s.ch //the prev character
	lit = string(s.ch)
	s.next()
	switch ch {
	case '(':
		fmt.Println("LPAREN")
	case ')':
		fmt.Println("RPAREN")
	case '=':
		fmt.Println("definition")
	case '"': //String
		return s.scanString()
	default:
		if s.offset >= len(s.src)-1 {
			fmt.Println("ende erreicht")
		} /*else {
			return s.scanIdentifier()
		}*/
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

/*
func (s *Scanner) scanIdentifier() (string, Pos) {
	start := s.offset

	for unicode.IsLetter(s.ch) || unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	lit := s.src[start:offset]
	return lit
}*/
/*
//Cosumes a number
func (s *Scanner) scanNumber() Node {
	start := s.offset

	for unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	return NewNumber(s.src[start:offset])
}*/

//Consumes one string
func (s *Scanner) scanString() Node {
	start := s.offset //start of the string
	sp := s.pos       //safe the position
	//TODO: implement strings correctly

	for s.ch != '"' {
		if s.ch == rune(0) {
			return String{s.src[start:s.offset], sp}
		}
		s.next()
		if s.ch == '\\' { //skipp escape sequences
			s.next()
		}
	}
	s.next()
	offset := s.offset //end of the string
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
	fmt.Println(s)
	sc := Scanner{}
	sc.Init(s)
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
	fmt.Println("Hallo")
	s := Symbol{"Hallo", Position{0, 1}}
	fmt.Println(s)
	f := &File{}
	f.Read("(blablub)laä#k(aslkdj)asldkjf($a $b=b) \"Hallo\"")
}

type Part interface{}
