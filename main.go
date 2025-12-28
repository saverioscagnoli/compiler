package main

import (
	"compiler/lexer"
	"fmt"
)

func main() {
	test := "var sdlfjsflx = 2323 + 3 * 4; var i = 0; while i < 10 { i = i + 1 }"
	l := lexer.New(test)

	for token := l.NextToken(); token.Type != lexer.Eof; token = l.NextToken() {
		fmt.Println(token)
	}
}
