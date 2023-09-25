package main

import (
	"fmt"

	"github.com/david-allan-jones/gomi-go-parser/frontend"
)

func main() {
	src := `
		# This is a test
		( （
		) ）
		? ？
		: ：
		, 、 ，
		[ 【
		] 】
		{ ｛
		} ｝
		! ！
		!= ！＝
		< ＜
		> ＞
		<= ＜＝
		>= ＞＝
		== ＝＝
		= ＝
		'abc' ”あいう”
		'a\aa'
		|| ｜｜
		&& ＆＆
		0 ０
		1 １
		0.1 ０．１
		10.0 １０．０
		abc
		あいう
		a1
		あ１
		a_1
		あ＿２
		+ ＋
		-
		* ＊
		/ ／
		% ％
		^ ＾
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
