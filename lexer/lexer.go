package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}
	return l
}

func (l *Lexer) readChar() {
	l.ch = l.input[l.readPosition]
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	var word Word

	for i := 0; i < len(l.input); i++ {
		l.readChar()
		var sep bool = isSeparetor(l.ch)

		if len(word.Value) > 0 || i == len(l.input)-1 {
			var sim bool = isSymbol(l.ch)
			var num bool = isNumber(l.ch)
			switch word.Type {
			case Symbol:
				if !sim || sep {
					tokens = append(tokens, decodeToken(word.Value))
					word = Word{}
				}
			case Number:
				if !num || sep {
					tokens = append(tokens, Token{Type: NUMBER, Literal: word.Value})
					word = Word{}
				}
			case String:
				if sim || sep {
					tokens = append(tokens, decodeToken(word.Value))
					word = Word{}
				}
			}
		}

		if !sep {
			word.append(l.ch)
		}
	}
	return tokens
}
