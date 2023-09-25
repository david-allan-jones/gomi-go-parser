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
		'a\\a'
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

	lexer := frontend.MakeGomiLexer(src)
	for token, err := lexer.ReadToken(); err == nil; token, err = lexer.ReadToken() {
		fmt.Println(token)
	}
}
