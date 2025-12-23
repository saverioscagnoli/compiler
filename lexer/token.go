package lexer

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	Plus   TokenType = "PLUS"
	Minus  TokenType = "MINUS"
	Star   TokenType = "STAR"
	Slash  TokenType = "SLASH"
	Assign TokenType = "ASSIGN"
	LParen TokenType = "LPAREN"
	RParen TokenType = "RPAREN"
	Var    TokenType = "VAR"
	Ident  TokenType = "IDENT"
	Int    TokenType = "INT"
)

type Token struct {
	Type  TokenType
	value string
}

func NewToken(t TokenType, value string) Token {
	return Token{
		Type:  t,
		value: value,
	}
}

var STATIC_TOKENS = map[string]TokenType{
	"+":   Plus,
	"-":   Minus,
	"*":   Star,
	"/":   Slash,
	"=":   Assign,
	"(":   LParen,
	")":   RParen,
	"var": Var,
}

type Char struct {
	Value    rune
	IsLetter bool
	IsNumber bool
	IsSymbol bool
	IsSpace  bool
}

func NewChar(value rune) Char {
	return Char{
		Value:    value,
		IsLetter: unicode.IsLetter(value),
		IsNumber: unicode.IsNumber(value),
		IsSymbol: unicode.IsSymbol(value),
		IsSpace:  unicode.IsSpace(value),
	}
}

// Check if characters are semantically coherent
// (e.g, letters with letters, numbers with numbers, symbols with symbols)
func (c *Char) CompareType(other Char) bool {
	return c.IsLetter == other.IsLetter && c.IsNumber == other.IsNumber && c.IsSymbol == other.IsSymbol
}

type Word struct {
	Chars    []Char
	IsNumber bool
}

func NewWord() Word {
	return Word{
		Chars:    []Char{},
		IsNumber: true,
	}
}

func (w *Word) AppendChar(ch Char) {
	w.Chars = append(w.Chars, ch)
}

func (w *Word) Length() int {
	return len(w.Chars)
}

func (w *Word) LastChar() Char {
	if len(w.Chars) == 0 {
		return Char{}
	}

	return w.Chars[len(w.Chars)-1]
}

func (w *Word) ToString() string {
	var sb strings.Builder

	for _, ch := range w.Chars {
		sb.WriteRune(ch.Value)
	}

	return sb.String()
}

func (w *Word) Reset() {
	w.Chars = []Char{}
	w.IsNumber = true
}
