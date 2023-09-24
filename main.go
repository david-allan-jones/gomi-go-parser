package main

import (
	"fmt"

	"github.com/david-allan-jones/gomi-go-parser/frontend"
)

func main() {
	src := `
		# This is a test
		#1
		(?:,)
		#3 # 4
		#5
	`
	lexer := frontend.MakeGomiLexer([]rune(src))
	for token, err := lexer.ReadToken(); token.Kind != frontend.EOFTokenKind; token, err = lexer.ReadToken() {
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(token)
	}
}
