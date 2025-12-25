package main

import (
	"compiler/lexer"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println("leone")
	l := lexer.NewLexer("var i2 = 2 + 2")
	fmt.Println(l.Tokenize())
}
