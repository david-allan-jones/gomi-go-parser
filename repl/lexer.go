package repl

import (
	"fmt"

	"github.com/david-allan-jones/gomi-go-parser/frontend"
	"github.com/fatih/color"
)

func StartLexerRepl() {
	lexInput := func(input string) {
		lexer := frontend.MakeGomiLexer(input)
		for {
			tok, err := lexer.ReadToken()
			if err != nil {
				color.HiRed(err.Error())
				break
			}
			color.HiYellow(fmt.Sprintf("{ %v, \"%v\", %v:%v }", tok.Kind, tok.Value, tok.Line, tok.Column))
			if tok.Kind == frontend.EOFTokenKind {
				break
			}
		}
	}
	startRepl(lexInput)
}
