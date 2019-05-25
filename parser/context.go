package parser

import (
	"fmt"
	"io"
)

type Location struct {
	FirstLine   int
	FirstColumn int
	LastLine    int
	LastColumn  int
}

type parseContext struct {
	*tokenizer

	// TODO(patrick): define parse tree
	err error
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

func Parse(reader io.Reader) (interface{}, error) {
	tokenizer, err := newTokenizer(reader)
	if err != nil {
		return nil, err
	}

	ctx := &parseContext{
		tokenizer: tokenizer,
	}
	qlParse(ctx)

	return nil, ctx.err
}

func init() {
	qlErrorVerbose = true
}
