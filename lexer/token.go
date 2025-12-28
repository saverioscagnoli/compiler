package lexer

type TokenType string

const (
	// Operators
	Plus       TokenType = "PLUS"
	Minus      TokenType = "MINUS"
	Star       TokenType = "STAR"
	Slash      TokenType = "SLASH"
	Modulo     TokenType = "MODULO"
	Not        TokenType = "NOT"
	And        TokenType = "AND"
	Or         TokenType = "OR"
	BitwiseAnd TokenType = "BITWISE_AND"
	BitwiseOr  TokenType = "BITWISE_OR"

	Assign TokenType = "ASSIGN"

	// Delimiters
	LParen  TokenType = "LPAREN"
	RParen  TokenType = "RPAREN"
	LSquare TokenType = "LSQUARE"
	RSquare TokenType = "RSQUARE"
	LBrace  TokenType = "LBRACE"
	RBrace  TokenType = "RBRACE"
	Comma   TokenType = "COMMA"
	Dot     TokenType = "DOT"
	Colon   TokenType = "COLON"

	// Keywords
	Var   TokenType = "VAR"
	Const TokenType = "CONST"
	If    TokenType = "IF"
	Else  TokenType = "ELSE"
	For   TokenType = "FOR"
	While TokenType = "WHILE"

	// Comparision
	Equals         TokenType = "EQUALS"
	NotEquals      TokenType = "NOT_EQUALS"
	Greater        TokenType = "GREATER"
	Less           TokenType = "LESS"
	GreaterOrEqual TokenType = "GREATER_OR_EQUAL"
	LessOrEqual    TokenType = "LESS_OR_EQUAL"

	// Variable, types, etc
	Ident  TokenType = "IDENT"
	True   TokenType = "TRUE"
	False  TokenType = "FALSE"
	Int    TokenType = "INT"
	Float  TokenType = "FLOAT"
	String TokenType = "STRING"

	// Function
	Function TokenType = "FUNCTION"
	Return   TokenType = "RETURN"

	// Special tokens
	Eof       TokenType = "EOF"
	Semicolon TokenType = "Semicolon"
	Illegal   TokenType = "ILLEGAL"
)

var STATIC_TOKENS = map[string]TokenType{
	"+":      Plus,
	"-":      Minus,
	"*":      Star,
	"/":      Slash,
	"%":      Modulo,
	"!":      Not,
	"&":      BitwiseAnd,
	"|":      BitwiseOr,
	"=":      Assign,
	"(":      LParen,
	")":      RParen,
	"[":      LSquare,
	"]":      RSquare,
	"{":      LBrace,
	"}":      RBrace,
	".":      Dot,
	":":      Colon,
	",":      Comma,
	">":      Greater,
	"<":      Less,
	"var":    Var,
	"const":  Const,
	"if":     If,
	"else":   Else,
	"for":    For,
	"while":  While,
	"true":   True,
	"false":  False,
	"fn":     Function,
	"return": Return,
	";":      Semicolon,
}

type Token struct {
	Type  TokenType
	Value string
}

func newToken(t TokenType, value string) Token {
	return Token{
		Type:  t,
		Value: value,
	}
}
