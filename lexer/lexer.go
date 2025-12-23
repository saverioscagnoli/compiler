package lexer

type Lexer struct {
	input string
}

func New(input string) *Lexer {
	return &Lexer{
		input: input,
	}
}

func AppendToken(tokens []Token, word *Word) []Token {
	wordString := word.ToString()

	if tokenType, exists := STATIC_TOKENS[wordString]; exists {
		token := NewToken(tokenType, wordString)
		tokens = append(tokens, token)
	} else if word.IsNumber {
		token := NewToken(Int, wordString)
		tokens = append(tokens, token)
	} else {
		token := NewToken(Ident, wordString)
		tokens = append(tokens, token)
	}

	word.Reset()
	return tokens
}

func (l *Lexer) Tokenize() []Token {
	tokens := []Token{}
	currentWord := NewWord()

	// Word separators:
	// 1 - Whitespace
	// 2 - Semantically different characters

	for i := 0; i < len(l.input); i++ {
		r := rune(l.input[i])
		ch := NewChar(r)

		if ch.IsSpace && currentWord.Length() > 0 {
			tokens = AppendToken(tokens, &currentWord)
			continue
		}

		if currentWord.Length() > 0 {
			prevCh := currentWord.LastChar()

			// Types changed, different token
			if !ch.CompareType(prevCh) {
				tokens = AppendToken(tokens, &currentWord)
			}
		}

		currentWord.AppendChar(ch)

		if currentWord.IsNumber {
			currentWord.IsNumber = ch.IsNumber
		}
	}

	// Append the last word if any
	if currentWord.Length() > 0 {
		tokens = AppendToken(tokens, &currentWord)
	}

	return tokens
}
