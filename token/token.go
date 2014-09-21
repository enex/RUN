package token

type Token int

//TODO: remove unused tokens

const (
	tok_start Token = iota

	EOF
	ILLEGAL
	COMMENT
	LBR //line break/ new line

	lit_start
	SYMBOL
	INTEGER
	FLOAT
	STRING
	lit_end

	op_start
	LPAREN
	RPAREN
	op_end

	key_start
	//keys
	key_end

	tok_end
)

var tok_strings = map[Token]string{
	EOF:     "EOF",
	ILLEGAL: "Illegal",
	COMMENT: "Comment",
	SYMBOL:  "Symbol",
	INTEGER: "Integer",
	LPAREN:  "(",
	RPAREN:  ")",
	STRING:  "string",
	FLOAT:   "float",
	LBR:     "lbr",
}

func (t Token) IsLiteral() bool {
	return t > lit_start && t < lit_end
}

func (t Token) IsOperator() bool {
	return t > op_start && t < op_end
}

func (t Token) IsKeyword() bool {
	return t > key_start && t < key_end
}

func Lookup(str string) Token {
	for t, s := range tok_strings {
		if s == str {
			return t
		}
	}
	return IDENT
}

func (t Token) String() string {
	return tok_strings[t]
}

func (t Token) Valid() bool {
	return t > tok_start && t < tok_end
}
