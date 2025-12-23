package main

import (
	"compiler/lexer"
	"fmt"
)

func main() {
	test := "var sdlfjsflx = 2323 + 3 * 4"
	lexer := lexer.New(test)

	tokens := lexer.Tokenize()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
