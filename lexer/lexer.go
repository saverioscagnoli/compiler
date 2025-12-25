package lexer

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

func decodeSymbol(word Word) Token {
	switch word.Value {
	case "+":
		return Token{Type: PLUS, Literal: "+"}
	case "-":
		return Token{Type: MINUS, Literal: "-"}
	case "*":
		return Token{Type: MULTIPLY, Literal: "*"}
	case "/":
		return Token{Type: DIVIDE, Literal: "/"}
	case "=":
		return Token{Type: EQUAL, Literal: "="}
	default:
		return Token{Type: ILLEGAL, Literal: word.Value}
	}
}

func decodeString(word Word) Token {
	return Token{Type: STRING, Literal: word.Value}
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	var word Word

	for i := 0; i < len(l.input); i++ {
		l.readChar()
		var sep bool = isSeparetor(l.ch)

		if len(word.Value) > 0 || i == len(l.input)-1 {
			var sim bool = isSymbol(l.ch)
			switch word.Type {
			case Symbol:
				if !sim || sep {
					tokens = append(tokens, decodeSymbol(word))
					word = Word{}
				}
			case Number:
				if !isNumber(l.ch) || sep {
					tokens = append(tokens, Token{Type: NUMBER, Literal: word.Value})
					word = Word{}
				}
			case String:
				if sim || sep {
					tokens = append(tokens, decodeString(word))
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
