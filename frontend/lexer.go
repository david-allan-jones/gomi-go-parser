package frontend

import (
	"fmt"
	"regexp"
)

type tokenKind string

const (
	ModuleTokenKind       tokenKind = "Module"
	ImportTokenKind       tokenKind = "Import"
	IntTokenKind          tokenKind = "Int"
	FloatTokenKind        tokenKind = "Float"
	BooleanTokenKind      tokenKind = "Boolean"
	StringTokenKind       tokenKind = "String"
	NilTokenKind          tokenKind = "Nil"
	IdentifierTokenKind   tokenKind = "Identifier"
	EqualsTokenKind       tokenKind = "Equals"
	OpenParenTokenKind    tokenKind = "OpenParen"
	CloseParenTokenKind   tokenKind = "CloseParen"
	BinOpTokenKind        tokenKind = "BinOp"
	OpenBracketTokenKind  tokenKind = "OpenBracket"
	CloseBracketTokenKind tokenKind = "CloseBracket"
	OpenBraceTokenKind    tokenKind = "OpenBrace"
	CloseBraceTokenKind   tokenKind = "CloseBrace"
	ColonTokenKind        tokenKind = "Colon"
	PeriodTokenKind       tokenKind = "Period"
	SemicolonTokenKind    tokenKind = "Semicolon"
	CommaTokenKind        tokenKind = "Comma"
	QuestionTokenKind     tokenKind = "Question"
	BangTokenKind         tokenKind = "Bang"
	LetTokenKind          tokenKind = "Let"
	ConstTokenKind        tokenKind = "Const"
	FuncTokenKind         tokenKind = "Func"
	IfTokenKind           tokenKind = "If"
	WhileTokenKind        tokenKind = "While"
	EOFTokenKind          tokenKind = "EOF"
)

var reserved = map[string]tokenKind{
	"module": ModuleTokenKind,
	"モジュール":  ModuleTokenKind,
	"import": ImportTokenKind,
	"インポート":  ImportTokenKind,
	"let":    LetTokenKind,
	"宣言":     LetTokenKind,
	"const":  ConstTokenKind,
	"定数":     ConstTokenKind,
	"nil":    NilTokenKind,
	"無":      NilTokenKind,
	"true":   BooleanTokenKind,
	"本当":     BooleanTokenKind,
	"false":  BooleanTokenKind,
	"嘘":      BooleanTokenKind,
	"if":     IfTokenKind,
	"もし":     IfTokenKind,
	"while":  WhileTokenKind,
	"繰り返す":   WhileTokenKind,
	"func":   FuncTokenKind,
	"関数":     FuncTokenKind,
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

func isDigit(ch rune) bool {
	sampleRegexp := regexp.MustCompile(`\d|[０-９]`)
	return sampleRegexp.MatchString(string(ch))
}

func identifierBeginAllowed(ch rune) bool {
	sampleRegexp := regexp.MustCompile(`[a-zA-Z\x{3041}-\x{3096}\x{30a1}-\x{30f6}\x{4e00}-\x{9faf}\x{30fc}]`)
	return sampleRegexp.MatchString(string(ch))
}

func identifierAllowed(ch rune) bool {
	return identifierBeginAllowed(ch) || isDigit(ch) || ch == '_' || ch == '＿'
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

func (lexer *gomiLexer) scanSingleCharToken() (token, error) {
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
		return lexer.makeToken(string(lexer.src[lexer.cursor]), CloseBracketTokenKind)
	}
	if lexer.src[lexer.cursor] == '{' || lexer.src[lexer.cursor] == '｛' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), OpenBraceTokenKind)
	}
	if lexer.src[lexer.cursor] == '}' || lexer.src[lexer.cursor] == '｝' {
		return lexer.makeToken(string(lexer.src[lexer.cursor]), CloseBraceTokenKind)
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
	return token{}, fmt.Errorf("no single char token found")
}

func (lexer *gomiLexer) scanArithBinOpToken() (token, error) {
	ch := lexer.at()
	if ch == '+' || ch == '＋' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	if ch == '-' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	if ch == '*' || ch == '＊' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	if ch == '/' || ch == '／' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	if ch == '%' || ch == '％' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	if ch == '^' || ch == '＾' {
		return lexer.makeToken(string(ch), BinOpTokenKind)
	}
	return token{}, fmt.Errorf("no arithmetic binop token")
}

func (lexer *gomiLexer) scanTwoCharCandidateToken(c1 rune, c2 rune, t1 tokenKind, t2 tokenKind) (token, error) {
	value := string(lexer.at())
	if lexer.at() == c1 {
		lexer.cursor++
		if lexer.at() == c2 {
			return lexer.makeToken(value+string(lexer.at()), t1)
		}
		lexer.cursor--
		return lexer.makeToken(value, t2)
	}
	return token{}, fmt.Errorf("could not scan two char token")
}

func (lexer *gomiLexer) scanBangEqualityToken() (token, error) {
	if tok, err := lexer.scanTwoCharCandidateToken('!', '=', BinOpTokenKind, BangTokenKind); err == nil {
		return tok, nil
	}
	if tok, err := lexer.scanTwoCharCandidateToken('！', '＝', BinOpTokenKind, BangTokenKind); err == nil {
		return tok, nil
	}
	return token{}, fmt.Errorf("no bang equality token")
}

func (lexer *gomiLexer) scanComparisonToken() (token, error) {
	if tok, err := lexer.scanTwoCharCandidateToken('<', '=', BinOpTokenKind, BinOpTokenKind); err == nil {
		return tok, nil
	}
	if tok, err := lexer.scanTwoCharCandidateToken('＜', '＝', BinOpTokenKind, BinOpTokenKind); err == nil {
		return tok, nil
	}
	if tok, err := lexer.scanTwoCharCandidateToken('>', '=', BinOpTokenKind, BinOpTokenKind); err == nil {
		return tok, nil
	}
	if tok, err := lexer.scanTwoCharCandidateToken('＞', '＝', BinOpTokenKind, BinOpTokenKind); err == nil {
		return tok, nil
	}
	return token{}, fmt.Errorf("no comparison token")
}

func (lexer *gomiLexer) scanEqualityToken() (token, error) {
	if tok, err := lexer.scanTwoCharCandidateToken('=', '=', BinOpTokenKind, EqualsTokenKind); err == nil {
		return tok, nil
	}
	if tok, err := lexer.scanTwoCharCandidateToken('＝', '＝', BinOpTokenKind, EqualsTokenKind); err == nil {
		return tok, nil
	}
	return token{}, fmt.Errorf("no equality token")
}

func (lexer *gomiLexer) scanStringToken() (token, error) {
	if lexer.at() == '\'' || lexer.at() == '”' {
		value := ""
		lexer.cursor++
		for lexer.at() != '\'' && lexer.at() != '”' {
			if lexer.at() == '\\' {
				lexer.cursor++
			}
			if lexer.at() == '\n' || lexer.cursor >= len(lexer.src) {
				return token{}, fmt.Errorf("strings must be closed and expressed on one line")
			}
			value += string(lexer.at())
			lexer.cursor++
		}
		return lexer.makeToken(value, StringTokenKind)
	}
	return token{}, fmt.Errorf("no string token")
}

func (lexer *gomiLexer) scanLogicalOpToken() (token, error) {
	if lexer.at() == '|' {
		lexer.cursor++
		if lexer.at() == '|' {
			return lexer.makeToken("||", BinOpTokenKind)
		}
	} else if lexer.at() == '｜' {
		lexer.cursor++
		if lexer.at() == '｜' {
			return lexer.makeToken("｜｜", BinOpTokenKind)
		}
	} else if lexer.at() == '&' {
		lexer.cursor++
		if lexer.at() == '&' {
			return lexer.makeToken("&&", BinOpTokenKind)
		}
	} else if lexer.at() == '＆' {
		lexer.cursor++
		if lexer.at() == '＆' {
			return lexer.makeToken("＆＆", BinOpTokenKind)
		}
	}
	return token{}, fmt.Errorf("no logical binop token")
}

func (lexer *gomiLexer) scanNumericToken() (token, error) {
	if isDigit(lexer.at()) {
		isFloat := false
		value := string(lexer.at())
		lexer.cursor++
		for lexer.cursor < len(lexer.src) && (isDigit(lexer.at()) || lexer.at() == '.' || lexer.at() == '．') {
			if lexer.at() == '.' || lexer.at() == '．' {
				isFloat = true
				value += string(lexer.at())
				lexer.cursor++
				if !isDigit(lexer.at()) {
					return token{}, fmt.Errorf("unexpected token while parsing float")
				}
				for lexer.cursor < len(lexer.src) && isDigit(lexer.at()) {
					value += string(lexer.at())
					lexer.cursor++
				}
			} else {
				value += string(lexer.at())
				lexer.cursor++
			}
		}
		lexer.cursor--
		if isFloat {
			return lexer.makeToken(value, FloatTokenKind)
		}
		return lexer.makeToken(value, IntTokenKind)
	}
	return token{}, fmt.Errorf("no numeric token")
}

func (lexer *gomiLexer) scanIdentifierToken() (token, error) {
	if identifierBeginAllowed(lexer.at()) {
		value := string(lexer.at())
		lexer.cursor++
		for lexer.cursor < len(lexer.src) && identifierAllowed(lexer.at()) {
			value += string(lexer.at())
			lexer.cursor++
		}
		lexer.cursor--
		if kind, has := reserved[value]; has {
			return lexer.makeToken(value, kind)
		}
		return lexer.makeToken(value, IdentifierTokenKind)
	}
	return token{}, fmt.Errorf("no identifier token")
}

func (lexer *gomiLexer) at() rune {
	return lexer.src[lexer.cursor]
}

func (lexer *gomiLexer) ReadToken() (token, error) {
	unrecognizedHit := false
	for !unrecognizedHit {
		lexer.eatSkippables()
		if lexer.cursor >= len(lexer.src) {
			return lexer.makeToken("EOF", EOFTokenKind)
		}
		if tok, err := lexer.scanSingleCharToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanArithBinOpToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanBangEqualityToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanComparisonToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanEqualityToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanStringToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanLogicalOpToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanNumericToken(); err == nil {
			return tok, nil
		}
		if tok, err := lexer.scanIdentifierToken(); err == nil {
			return tok, nil
		}
		unrecognizedHit = true
	}
	return token{}, fmt.Errorf(`
		Lexer Error!
		Unrecognized token: '%v'
		Line: %v
		Column: %v
	`, string(lexer.src[lexer.cursor]), lexer.line, lexer.column)
}
