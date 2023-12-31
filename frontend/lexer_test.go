package frontend

import (
	"testing"

	"github.com/david-allan-jones/gomi-go-parser/frontend/errors"
)

func testSources(t *testing.T, programs []string, kind tokenKind) {
	for i := 0; i < len(programs); i++ {
		lexer := MakeGomiLexer(programs[i])
		tok, _ := lexer.ReadToken()
		if tok.Value != programs[i] {
			t.Fatalf("Failed value test. Received %v but expected %v", tok.Value, programs[i])
		}
		if tok.Kind != kind {
			t.Fatalf("Failed kind test. Received '%v' but expected '%v'", tok.Kind, kind)
		}
	}
}

func TestEmptyString(t *testing.T) {
	lexer := MakeGomiLexer("")
	tok, _ := lexer.ReadToken()
	if tok.Kind != EOFTokenKind {
		t.Fail()
	}
}

func TestLineComment(t *testing.T) {
	lexer := MakeGomiLexer(`
		# This is a comment
		1	
	`)
	tok, _ := lexer.ReadToken()
	if tok.Kind != IntTokenKind || tok.Value != "1" {
		t.Fail()
	}
}

func TestOpenParen(t *testing.T) {
	testSources(t, []string{"(", "（"}, OpenParenTokenKind)
}

func TestCloseParen(t *testing.T) {
	testSources(t, []string{")", "）"}, CloseParenTokenKind)
}

func TestOpenBracket(t *testing.T) {
	testSources(t, []string{"[", "【"}, OpenBracketTokenKind)
}

func TestCloseBracket(t *testing.T) {
	testSources(t, []string{"]", "】"}, CloseBracketTokenKind)
}

func TestOpenBrace(t *testing.T) {
	testSources(t, []string{"{", "｛"}, OpenBraceTokenKind)
}

func TestCloseBrace(t *testing.T) {
	testSources(t, []string{"}", "｝"}, CloseBraceTokenKind)
}

func TestColon(t *testing.T) {
	testSources(t, []string{":", "："}, ColonTokenKind)
}

func TestSemicolon(t *testing.T) {
	testSources(t, []string{";", "；"}, SemicolonTokenKind)
}

func TestComma(t *testing.T) {
	testSources(t, []string{",", "、", "，"}, CommaTokenKind)
}

func TestQuestion(t *testing.T) {
	testSources(t, []string{"?", "？"}, QuestionTokenKind)
}

func TestBang(t *testing.T) {
	testSources(t, []string{"!", "！"}, BangTokenKind)
}

func TestInt(t *testing.T) {
	testSources(t, []string{"1", "１", "10", "１０"}, IntTokenKind)
}

func TestFloat(t *testing.T) {
	testSources(t, []string{"1.0", "１．０", "10.01", "１０．０１"}, FloatTokenKind)
}

func TestString(t *testing.T) {
	tests := []string{"''", "'Test'", "””", "”テスト”"}
	for i := 0; i < len(tests); i++ {
		lexer := MakeGomiLexer(tests[i])
		tok, _ := lexer.ReadToken()
		if tok.Kind != StringTokenKind {
			t.Fatalf("Failed to lex a string. Received %v", tok.Kind)
		}
	}
}

func TestNil(t *testing.T) {
	testSources(t, []string{"nil", "無"}, NilTokenKind)
}

func TestBoolean(t *testing.T) {
	testSources(t, []string{"true", "本当", "false", "嘘"}, BooleanTokenKind)
}

func TestIf(t *testing.T) {
	testSources(t, []string{"if", "もし"}, IfTokenKind)
}

func TestWhile(t *testing.T) {
	testSources(t, []string{"while", "繰り返す"}, WhileTokenKind)
}

func TestLet(t *testing.T) {
	testSources(t, []string{"let", "宣言"}, LetTokenKind)
}

func TestConst(t *testing.T) {
	testSources(t, []string{"const", "定数"}, ConstTokenKind)
}

func TestModule(t *testing.T) {
	testSources(t, []string{"module", "モジュール"}, ModuleTokenKind)
}

func TestImport(t *testing.T) {
	testSources(t, []string{"import", "インポート"}, ImportTokenKind)
}

func TestFunc(t *testing.T) {
	testSources(t, []string{"func", "関数"}, FuncTokenKind)
}

func TestIdentifier(t *testing.T) {
	testSources(t, []string{"a", "a1", "あ", "あ１", "a_x"}, IdentifierTokenKind)
}

func TestAssignment(t *testing.T) {
	testSources(t, []string{"=", "＝"}, EqualsTokenKind)
}

func TestSingleCharBinOps(t *testing.T) {
	testSources(t, []string{"+", "-", "*", "/", "%", "^", "＋", "＊", "／", "％", "＾"}, BinOpTokenKind)
}

func TestMultipleCharBinOps(t *testing.T) {
	testSources(t, []string{"||", "&&", "==", "!=", "<=", ">=", "｜｜", "＆＆", "＝＝", "！＝", "＜＝", "＞＝"}, BinOpTokenKind)
}

func TestWhitespaceAtEndOfFile(t *testing.T) {
	var toks []token
	lexer := MakeGomiLexer(`
		a 
	`)
	for tok, err := lexer.ReadToken(); err == nil; tok, err = lexer.ReadToken() {
		toks = append(toks, tok)
	}
	// Add 1 for EOF
	if len(toks) != 2 {
		t.Fatalf("Expected 2 tokens but got %v", len(toks))
	}
}

func TestUnrecognizedChars(t *testing.T) {
	lexer := MakeGomiLexer("$")
	_, err := lexer.ReadToken()
	if err == nil {
		t.Fatalf("Expected error on unrecognized character but none was thrown")
	}
	if err.Kind != errors.UnrecognizedError {
		t.Fatalf("Unrecognized char errored but it was of type %v", err.Kind)
	}
}

func TestEofError(t *testing.T) {
	lexer := MakeGomiLexer("")
	lexer.ReadToken()
	_, err := lexer.ReadToken()
	if err == nil {
		t.Fatalf("Expected error on unrecognized character but none was thrown")
	}
	if err.Kind != errors.EofError {
		t.Fatalf("Unrecognized char errored but it was of type %v", err.Kind)
	}
}

func TestLineNumberTracking(t *testing.T) {
	lexer := MakeGomiLexer(`
		a	
	`)
	tok, _ := lexer.ReadToken()
	if tok.Line != 2 {
		t.Fail()
	}
}

func TestColumnTracking(t *testing.T) {
	lexer := MakeGomiLexer(" ( ")
	tok, _ := lexer.ReadToken()
	if tok.Column != 2 {
		t.Fail()
	}
}
