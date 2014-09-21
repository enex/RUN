package scan

import (
	"fmt"
	"github.com/enex/RUN/ast"
	"github.com/enex/RUN/token"
	"unicode"
)

// Scanner
type Scanner struct {
	ch      rune //last charracter scanned
	offset  int  //offset of currnt character
	roffset int  //offset of next character

	pos token.Position

	src string //actual source code

	file  *token.File
	ident int
}

// Init initializes Scanner and makes the source code ready to Scan
func (s *Scanner) Init(src string) {
	s.offset, s.roffset = 0, 0
	s.src = src

	s.next()
}

func New(src string) Scanner {
	s := Scanner{}
	s.Init(src)
	return s
}

func (s *Scanner) scan() ast.Node {
	s.skipWhitespace()
	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}

	//fmt.Println("scan", string(s.ch))

	switch s.ch {
	case '(':
		return s.scanParen()
	case ')':
		fmt.Println("RPAREN")
		//panic("a paren has been closed without an opening paren")
	/*case '=':
	fmt.Println("definition")*/
	case '"': //String
		return s.scanString()
	case '/':
		v := s.ch
		p, pr := s.offset, s.roffset
		s.next()
		if s.ch == '/' {
			s.consumeComment()
		} else {
			s.ch = v //set it back to the previous value
			s.offset = p
			s.roffset = pr
		}
	case '\n':
		fmt.Println("Zeilenumbruch")
	default:
		if s.offset >= len(s.src)-1 {
			fmt.Println("ende erreicht")
			return nil
		}
		fmt.Println("irgend was anderes")
	}
	s.next()
	return nil
}

func (s *Scanner) Scan() ast.Node {
	scope := make([]ast.Node, 0)
	for !(s.offset >= len(s.src)-1) {
		n := s.scanLine()
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
			s.pos.Row++
			s.pos.Col = 0
		} else {
			s.pos.Col++
		}
		s.roffset++
	}
}

func (s *Scanner) scanLine() ast.Node {
	startIdent := s.ident //Save the ident when beginning scanning
	e := ast.Paren{Position: s.pos, Content: make([]ast.Node, 0)}
	for {
		e.Content = append(e.Content, s.scan())
		//fmt.Println(e)
		if s.ch == '\n' {
			//fmt.Println("new line")
			i := 0
			s.next()
			for s.ch == '\t' {
				i++
				//fmt.Println("ident", i)
				s.next()
			}
			s.ident = i
			//fmt.Println(i, startIdent)
			for s.ident > startIdent {
				e.Content = append(e.Content, s.scanLine())
			}
			break
			/*
				offset := s.offset
				roffset := s.roffset

				s.offset = offset
				roffset = roffset
			*/
		}
		if s.offset >= len(s.src)-1 {
			//fmt.Println("ende erreicht")
			break
		}
	}
	return e
}

func (s *Scanner) scanParen() ast.Node {
	sp := s.pos //safe the position
	s.next()    //go to the next char
	content := make([]ast.Node, 0)
	for s.ch != ')' {
		//fmt.Println("consume")
		content = append(content, s.scan())
	}
	s.next()
	return ast.Paren{Position: sp, Content: content}
}

//Cosumes a number
func (s *Scanner) scanNumber() ast.Node {
	start := s.offset
	sp := s.pos
	for unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	//fmt.Println("scan Number")
	return ast.NewNumber(s.src[start:offset], sp)
}

//Consumes one string
func (s *Scanner) scanString() ast.Node {
	sp := s.pos //safe the position
	s.next()
	r := ""
	//TODO: implement strings correctly
	for s.ch != '"' {
		r += string(s.ch)
		if s.ch == rune(0) {
			break
		}
		s.next()
		if s.ch == '\\' { //skipp escape sequences
			s.next()
			r += string(s.ch)
			s.next()
		}
	}
	s.next() //skipp the las "
	return ast.NewString(r, sp)
}

//consume a Comment, it will be thrown away afterwards
func (s *Scanner) consumeComment() {
	for s.ch != '\n' && s.offset < len(s.src)-1 {
		s.next()
	}
	s.next()
}

//will skipp all the white spaces without caring about them
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' {
		s.next()
	}
}
