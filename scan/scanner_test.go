package scan

import (
	"github.com/enex/RUN/token"
	"testing"
)

//function to simplify testing for correctness
func test_handler(t *testing.T, src string, expected []token.Token) {
	var s Scanner
	s.Init(token.NewFile("", 1, len(src)), src)
	lit, tok, pos := s.Scan()
	for i := 0; tok != token.EOF; i++ {
		t.Log("tok:", tok.String(), "|", lit, "|", pos)
		if i >= len(expected) {
			t.Fatal(pos, "to meny tokens")
		}
		if tok != expected[i] {
			t.Fatal(pos, "Expected:", expected[i], "Got:", tok)
		}
		lit, tok, pos = s.Scan()
	}
}

func TestNumber(t *testing.T) {
	test_handler(t, "9", []token.Token{
		token.INTEGER,
		token.EOF,
	})
}

func TestString(t *testing.T) {
	test_handler(t, "\"Hello World\"", []token.Token{
		token.STRING,
		token.EOF,
	})
}

func TestComment(t *testing.T) {
	test_handler(t, "234//Hallo", []token.Token{
		token.INTEGER,
		token.COMMENT,
		token.EOF,
	})
}

//test how good paren works with ()
func TestLParen(t *testing.T) {
	test_handler(t, "(a + b)(a - c (23 + 6))", []token.Token{
		token.LPAREN,
		token.SYMBOL,
		token.SYMBOL,
		token.SYMBOL,
		token.RPAREN,
		token.LPAREN,
		token.SYMBOL,
		token.SYMBOL,
		token.SYMBOL,
		token.LPAREN,
		token.INTEGER,
		token.SYMBOL,
		token.INTEGER,
		token.RPAREN,
		token.RPAREN,
		token.EOF,
	})
}
