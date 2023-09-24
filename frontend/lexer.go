package frontend

import "fmt"

type tokenValue string

const (
	EN_MODULE        tokenValue = "module"
	JP_MODULE        tokenValue = "モジュール"
	EN_IMPORT        tokenValue = "import"
	JP_IMPORT        tokenValue = "インポート"
	HW_EQUALS        tokenValue = "="
	FW_EQUALS        tokenValue = "＝"
	HW_OPEN_PAREN    tokenValue = "("
	FW_OPEN_PAREN    tokenValue = "（"
	HW_CLOSE_PAREN   tokenValue = ")"
	FW_CLOSE_PAREN   tokenValue = "）"
	HW_OPEN_BRACKET  tokenValue = "["
	FW_OPEN_BRACKET  tokenValue = "【"
	HW_CLOSE_BRACKET tokenValue = "]"
	FW_CLOSE_BRACKET tokenValue = "】"
	HW_OPEN_BRACE    tokenValue = "{"
	FW_OPEN_BRACE    tokenValue = "｛"
	HW_CLOSE_BRACE   tokenValue = "}"
	FW_CLOSE_BRACE   tokenValue = "｝"
	HW_COLON         tokenValue = ":"
	FW_COLON         tokenValue = "："
	HW_PERIOD        tokenValue = "."
	FW_PERIOD        tokenValue = "。"
	HW_SEMICOLON     tokenValue = ";"
	FW_SEMICOLON     tokenValue = "；"
	HW_COMMA         tokenValue = ","
	FW_COMMA_1       tokenValue = "，"
	FW_COMMA_2       tokenValue = "、"
	HW_QUESTION      tokenValue = "?"
	FW_QUESTION      tokenValue = "？"
	HW_BANG          tokenValue = "!"
	FW_BANG          tokenValue = "！"
	EN_COMMENT       tokenValue = "#"
	JP_COMMENT       tokenValue = "＃"
	EN_STRING        tokenValue = "'"
	JP_STRING        tokenValue = "”"
	EN_FUNC          tokenValue = "func"
	JP_FUNC          tokenValue = "関数"
	EN_WHILE         tokenValue = "while"
	JP_WHILE         tokenValue = "繰り返す"
	EN_LET           tokenValue = "let"
	JP_LET           tokenValue = "宣言"
	EN_CONST         tokenValue = "const"
	JP_CONST         tokenValue = "定数"
	EN_NIL           tokenValue = "nil"
	JP_NIL           tokenValue = "無"
	EN_TRUE          tokenValue = "true"
	JP_TRUE          tokenValue = "本当"
	EN_FALSE         tokenValue = "false"
	JP_FALSE         tokenValue = "嘘"
	EN_IF            tokenValue = "if"
	JP_IF            tokenValue = "もし"
	EOF              tokenValue = "EOF"
)

type tokenKind uint

const (
	ModuleTokenKind       tokenKind = 1
	ImportTokenKind       tokenKind = 2
	IntTokenKind          tokenKind = 3
	FloatTokenKind        tokenKind = 4
	BooleanTokenKind      tokenKind = 5
	StringTokenKind       tokenKind = 6
	NilTokenKind          tokenKind = 7
	IdentifierTokenKind   tokenKind = 8
	EqualsTokenKind       tokenKind = 9
	OpenParenTokenKind    tokenKind = 10
	CloseParenTokenKind   tokenKind = 11
	BinOpTokenKind        tokenKind = 12
	OpenBracketTokenKind  tokenKind = 13
	CloseBracketTokenKind tokenKind = 14
	OpenBraceTokenKind    tokenKind = 15
	CloseBraceTokenKind   tokenKind = 16
	ColonTokenKind        tokenKind = 17
	PeriodTokenKind       tokenKind = 18
	SemicolonTokenKind    tokenKind = 19
	CommaTokenKind        tokenKind = 20
	QuestionTokenKind     tokenKind = 21
	BangTokenKind         tokenKind = 22
	LetTokenKind          tokenKind = 23
	ConstTokenKind        tokenKind = 24
	FuncTokenKind         tokenKind = 25
	IfTokenKind           tokenKind = 26
	WhileTokenKind        tokenKind = 27
	EOFTokenKind          tokenKind = 28
)

var reserved = map[tokenValue]tokenKind{
	EN_MODULE: ModuleTokenKind,
	JP_MODULE: ModuleTokenKind,
	EN_IMPORT: ImportTokenKind,
	JP_IMPORT: ImportTokenKind,
	EN_LET:    LetTokenKind,
	JP_LET:    LetTokenKind,
	EN_CONST:  ConstTokenKind,
	JP_CONST:  ConstTokenKind,
	EN_NIL:    NilTokenKind,
	JP_NIL:    NilTokenKind,
	EN_TRUE:   BooleanTokenKind,
	JP_TRUE:   BooleanTokenKind,
	EN_FALSE:  BooleanTokenKind,
	JP_FALSE:  BooleanTokenKind,
	EN_IF:     IfTokenKind,
	JP_IF:     IfTokenKind,
	EN_WHILE:  WhileTokenKind,
	JP_WHILE:  WhileTokenKind,
	EN_FUNC:   FuncTokenKind,
	JP_FUNC:   FuncTokenKind,
}

type token struct {
	Value  string
	Kind   tokenKind
	Line   int
	Column int
}

type gomiLexer struct {
	src    []rune
	cursor int
	line   int
	column int
}

func MakeGomiLexer(src []rune) gomiLexer {
	return gomiLexer{
		src:    src,
		cursor: 0,
		line:   0,
		column: 0,
	}
}

func skippable(ch rune) bool {
	return ch == ' ' || ch == '　' || ch == '\n' || ch == '\r' || ch == '\t' || ch == '#' || ch == '＃'
}

func (lexer *gomiLexer) makeToken(value string, kind tokenKind) (token, error) {
	tok := token{
		Value:  value,
		Kind:   kind,
		Line:   lexer.line,
		Column: lexer.column,
	}
	if kind != EOFTokenKind {
		lexer.cursor++
		lexer.column += len([]rune(value))
	}
	return tok, nil
}

func (lexer *gomiLexer) eatSkippables() {
	inComment := false
	for lexer.cursor < len(lexer.src) && (skippable(lexer.src[lexer.cursor]) || inComment) {
		if lexer.src[lexer.cursor] == '#' || lexer.src[lexer.cursor] == '＃' {
			lexer.column++
			inComment = true
		} else if lexer.src[lexer.cursor] == '\n' {
			lexer.line++
			lexer.column = 1
			inComment = false
		} else {
			lexer.column++
		}
		lexer.cursor++
	}
}

func (lexer *gomiLexer) readSingleCharToken() (token, error) {
	if lexer.src[lexer.cursor] == '(' || lexer.src[lexer.cursor] == '（' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenParenTokenKind)
	}
	if lexer.src[lexer.cursor] == ')' || lexer.src[lexer.cursor] == '）' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), CloseParenTokenKind)
	}
	if lexer.src[lexer.cursor] == '[' || lexer.src[lexer.cursor] == '【' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenBracketTokenKind)
	}
	if lexer.src[lexer.cursor] == ']' || lexer.src[lexer.cursor] == '】' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenBracketTokenKind)
	}
	if lexer.src[lexer.cursor] == '{' || lexer.src[lexer.cursor] == '｛' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenBraceTokenKind)
	}
	if lexer.src[lexer.cursor] == '}' || lexer.src[lexer.cursor] == '｝' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenBraceTokenKind)
	}
	if lexer.src[lexer.cursor] == '.' || lexer.src[lexer.cursor] == '。' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), PeriodTokenKind)
	}
	if lexer.src[lexer.cursor] == ':' || lexer.src[lexer.cursor] == '：' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), ColonTokenKind)
	}
	if lexer.src[lexer.cursor] == ';' || lexer.src[lexer.cursor] == '；' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), SemicolonTokenKind)
	}
	if lexer.src[lexer.cursor] == ',' || lexer.src[lexer.cursor] == '、' || lexer.src[lexer.cursor] == '，' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), CommaTokenKind)
	}
	if lexer.src[lexer.cursor] == '?' || lexer.src[lexer.cursor] == '？' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), QuestionTokenKind)
	}
	return token{}, fmt.Errorf("No single char token found")
}

func (lexer *gomiLexer) ReadToken() (token, error) {
	for {
		lexer.eatSkippables()
		if lexer.cursor >= len(lexer.src) {
			return lexer.makeToken("EOF", EOFTokenKind)
		}
		tok, err := lexer.readSingleCharToken()
		if err == nil {
			return tok, nil
		}
		// tok, err = lexer.readBangEqualityToken()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readComparisonToken()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readEquality()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readStringToken()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readBinOpToken()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readNumericToken()
		// if err == nil {
		// 	return tok, nil
		// }
		// tok, err = lexer.readIdentiferToken()
		// if err == nil {
		// 	return tok, nil
		// }
		return token{}, fmt.Errorf(`
			Lexer Error!
			Unrecognized token: '%v'
			Line: %v
			Column: %v
		`, string(lexer.src[lexer.cursor]), lexer.line, lexer.column)
	}
}
