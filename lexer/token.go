package lexer

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType int

const (
	ILLEGAL = iota
	EOF
	STRING
	NUMBER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	EQUAL
)
