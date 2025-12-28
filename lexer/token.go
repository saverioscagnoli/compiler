package lexer

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"
	STRING  TokenType = "STRING"
	NUMBER  TokenType = "NUMBER"
	PLUS    TokenType = "PLUS"
	MINUS   TokenType = "MINUS"
	STAR    TokenType = "STAR"
	SLASH   TokenType = "SLASH"
	ASSIGN  TokenType = "ASSIGN"
	VAR     TokenType = "VAR"
)

type Word struct {
	Value string
	Type  WordType
}

type WordType int

const (
	Number WordType = iota
	String
	Symbol
)

func (w *Word) append(ch byte) {
	w.Value += string(ch)
	if isSymbol(ch) {
		w.Type = Symbol
	} else if !isNumber(ch) {
		w.Type = String
	}
}

func isNumber(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isSymbol(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '='
}

func isSeparetor(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

var STATIC_TOKENS = map[string]TokenType{
	"+":   PLUS,
	"-":   MINUS,
	"*":   STAR,
	"/":   SLASH,
	"=":   ASSIGN,
	"var": VAR,
}

func decodeToken(s string) Token {
	if tokenType, exists := STATIC_TOKENS[s]; exists {
		return Token{Type: tokenType, Literal: s}
	}
	return Token{Type: STRING, Literal: s}
}
