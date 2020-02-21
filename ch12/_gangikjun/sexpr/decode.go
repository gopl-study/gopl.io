package sexpr

import (
	"fmt"
	"go/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune // current token
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want { // NOTE: 좋은 예외 처리 예제가 아님
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}
