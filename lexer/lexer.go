package lexer

import (
	"unicode"
)

type Lexer struct {
	input   string
	current rune
	cursor  int
}

func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		cursor: 0,
	}

	// Initialize the current character
	l.advance()

	return l
}

func (l *Lexer) getCharAtCursor() rune {
	return rune(l.input[l.cursor])
}

func (l *Lexer) advance() {
	if l.cursor >= len(l.input) {
		l.current = 0
	} else {
		l.current = l.getCharAtCursor()
	}

	l.cursor++
}

func (l *Lexer) peek() rune {
	if l.cursor < len(l.input) {
		return rune(l.input[l.cursor])
	}

	return 0
}

func (l *Lexer) skipWhitespace() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		l.advance()
	}
}

// If a letter is found, keep reading until a non-alphanumeric character is encountered.
func (l *Lexer) readIdent() string {
	var ident string

	for unicode.IsLetter(l.current) || unicode.IsDigit(l.current) || l.current == '_' {
		ident += string(l.current)
		l.advance()
	}

	return ident
}

func (l *Lexer) readString() string {
	var literal string

	// Skip the opening quote
	l.advance()

	for l.current != '"' && l.current != 0 {
		// Handle escape sequences
		if l.current == '\\' {
			l.advance() // Skip the backslash

			switch l.current {
			case 'n':
				literal += "\n"
			case 't':
				literal += "\t"
			case 'r':
				literal += "\r"
			case '\\':
				literal += "\\"
			case '"':
				literal += "\""
			case '\'':
				literal += "'"
			case '0':
				literal += "\x00"
			default:
				// If unknown escape sequence, just include the character as-is
				literal += string(l.current)
			}

			l.advance()
		} else {
			literal += string(l.current)
			l.advance()
		}
	}

	// l.current here is '"', skip
	l.advance()

	return literal
}

// If a digit is found, keep reading until a non-digit character is encountered.
// If a decimal is found, also return true
func (l *Lexer) readNumber() (string, bool) {
	var number string
	var decimalFound bool

	for unicode.IsDigit(l.current) || (l.current == '.' && !decimalFound && unicode.IsDigit(l.peek())) {
		if l.current == '.' {
			decimalFound = true
		}

		number += string(l.current)
		l.advance()
	}

	return number, decimalFound
}

func (l *Lexer) NextToken() Token {
	var token Token

	l.skipWhitespace()

	if l.current == 0 {
		return newToken(Eof, "")
	}

	// Check single characters that could possibly part of a multi-character token "=="
	switch l.current {
	case '=':
		// Check for potential ==
		if l.peek() == '=' {
			token = newToken(Equals, "==")

			l.advance() // Skip this '='
			l.advance() // Skip next '='

			return token
		}

	case '!':
		// Check for potential !=
		if l.peek() == '=' {
			token = newToken(NotEquals, "!=")

			l.advance() // Skip this '!'
			l.advance() // Skip next '='

			return token
		}

	case '<':
		// Check for potential <=
		if l.peek() == '=' {
			token = newToken(LessOrEqual, "<=")

			l.advance() // Skip this '<'
			l.advance() // Skip next '='

			return token
		}

	case '>':
		// Check for potential >=
		if l.peek() == '=' {
			token = newToken(GreaterOrEqual, ">=")

			l.advance() // Skip this '>'
			l.advance() // Skip next '='

			return token
		}

	case '&':
		// Check for potential &&
		if l.peek() == '&' {
			token = newToken(And, "&&")

			l.advance() // Skip this '&'
			l.advance() // Skip next '&'

			return token
		}

	case '|':
		if l.peek() == '|' {
			token = newToken(Or, "||")

			l.advance() // Skip this '|'
			l.advance() // Skip next '|'

			return token
		}
	}

	// Check if the current character is a known single-character token "(, *" etc.
	if tokenType, ok := STATIC_TOKENS[string(l.current)]; ok {
		token = newToken(tokenType, string(l.current))
		l.advance()

		return token
	}

	switch {
	case l.current == '"':
		token = newToken(String, l.readString())

	case unicode.IsLetter(l.current) || l.current == '_':
		ident := l.readIdent()

		// First check if it's actually a known multi-character token "var, for, while" etc.
		if tokenType, ok := STATIC_TOKENS[ident]; ok {
			token = newToken(tokenType, ident)
		} else {
			token = newToken(Ident, ident)
		}

	case unicode.IsDigit(l.current):
		literal, isFloat := l.readNumber()

		if isFloat {
			token = newToken(Float, literal)
		} else {
			token = newToken(Int, literal)
		}

	default:
		token = newToken(Illegal, string(l.current))
	}

	return token
}
