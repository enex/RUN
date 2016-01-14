package scan

import (
	"github.com/enex/RUN/token"
	"unicode"
)

// Scanner
type Scanner struct {
	ch      rune //last charracter scanned
	offset  int  //offset of currnt character
	roffset int  //offset of next character

	src string //actual source code

	file *token.File

	ident int //last read ident in tabs
}

// Init initializes Scanner and makes the source code ready to Scan
func (s *Scanner) Init(file *token.File, src string) {
	s.file = file
	s.offset, s.roffset = 0, 0
	s.src = src
	s.file.AddLine(s.offset)

	s.next()
}

func New(file *token.File, src string) Scanner {
	s := Scanner{}
	s.Init(file, src)
	return s
}

func (s *Scanner) Scan() (lit string, tok token.Token, pos token.Pos) {
	s.skipWhitespace()

	if s.ch == '\n' { //Parse line break and ident after linebreak
		lit, pos = string(s.ch), s.file.Pos(s.offset)
		tok = token.LBR
		s.next()
		for s.ch == '\t' {
			lit += string(s.ch)
			s.next()
		}
		return
	}

	if unicode.IsLetter(s.ch) {
		return s.scanSymbol()
	}

	if unicode.IsDigit(s.ch) {
		return s.scanNumber()
	}

	ch := s.ch //the prev character
	lit, pos = string(s.ch), s.file.Pos(s.offset)
	s.next()
	switch ch {
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '/':
		if s.ch == '/' {
			s.next()
			return s.scanComment()
		}
	case '!': //a token directly leading to an illegal token
		tok = token.ILLEGAL
	case '"': //String
		return s.scanString()
	default:
		if s.offset >= len(s.src)-1 {
			tok = token.EOF
		} else {
			return s.scanSymbol()
		}
	}

	return
}

func (s *Scanner) next() {
	s.ch = rune(0)
	if s.roffset < len(s.src) {
		s.offset = s.roffset
		s.ch = rune(s.src[s.offset])
		if s.ch == '\n' {
			s.file.AddLine(s.offset)
		}
		s.roffset++
	}
}

func (s *Scanner) scanSymbol() (string, token.Token, token.Pos) {
	start := s.offset

	for unicode.IsLetter(s.ch) || unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	lit := s.src[start:offset]
	return lit, token.Lookup(lit), s.file.Pos(start)
}

//Cosumes a number
func (s *Scanner) scanNumber() (string, token.Token, token.Pos) {
	start := s.offset

	for unicode.IsDigit(s.ch) {
		s.next()
	}
	offset := s.offset
	if s.ch == rune(0) {
		offset++
	}
	return s.src[start:offset], token.INTEGER, s.file.Pos(start)
}

//Consumes one string
func (s *Scanner) scanString() (string, token.Token, token.Pos) {
	start := s.offset //start of the string

	//TODO: implement strings correctly

	for s.ch != '"' {
		if s.ch == rune(0) {
			return s.src[start:s.offset], token.ILLEGAL, s.file.Pos(start)
		}
		s.next() /*
			if s.ch == '\\' { //skipp escape sequences
				s.next()
			}*/
	}
	s.next()
	offset := s.offset //end of the string
	return s.src[start:offset], token.STRING, s.file.Pos(start)
}

//scanns a comment, comments can actually be used by the compiler, so it
//will be scanned and saved but maybe it will be thrown away later in the
//Programm because no one cares about it
func (s *Scanner) scanComment() (string, token.Token, token.Pos) {
	start := s.offset //start tof the comment
	for s.ch != '\n' && s.offset < len(s.src)-1 {
		s.next()
	}
	s.next()
	offset := s.offset + 1 //end of the comment
	return s.src[start:offset], token.COMMENT, s.file.Pos(start)
}

//will skipp all the white spaces without caring about them
func (s *Scanner) skipWhitespace() {
	for s.ch == ' ' || s.ch == '\t' {
		s.next()
	}
}
