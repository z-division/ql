package parser

import (
	"fmt"
	"io"
)

type parseContext struct {
	*tokenizer

	// TODO(patrick): define parse tree
	parsed []Expr

	err error
}

func (ctx *parseContext) setParsed(parsed []Expr) {
	ctx.parsed = parsed
}

// TODO(patrick): implement
func (ctx *parseContext) Lex(lval *qlSymType) int {
	token, err := ctx.nextToken(lval)
	if err != nil {
		ctx.err = err
		return LEX_ERROR
	}

	return token
}

func (ctx *parseContext) Error(msg string) {
	if ctx.err == nil {
		// TODO(patrick): include location info from lexer
		ctx.err = fmt.Errorf(msg)
	}
}

func Parse(filename string, reader io.Reader) ([]Expr, error) {
	tokenizer, err := newTokenizer(filename, reader)
	if err != nil {
		return nil, err
	}

	ctx := &parseContext{
		tokenizer: tokenizer,
	}
	qlParse(ctx)

	return ctx.parsed, ctx.err
}

func init() {
	qlErrorVerbose = true
}
