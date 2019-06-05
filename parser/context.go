package parser

import (
	"fmt"
	"io"
)

type parseContext struct {
	// TODO(patrick): switch to yet to be written tokenizer
	*rawTokenizer

	// TODO(patrick): define parse tree
	parsed []Node

	err error
}

func (ctx *parseContext) setParsed(parsed []Node) {
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
		pos := Location{
			Filename: ctx.filename,
			Start:    Position{1, 1},
			End:      Position{1, 1},
		}
		if ctx.prevToken != nil {
			pos = ctx.prevToken.Location
		}
		ctx.err = fmt.Errorf("%v: %s", pos, msg)
	}
}

func Parse(filename string, reader io.Reader) ([]Node, error) {
	tokenizer, err := newRawTokenizer(filename, reader)
	if err != nil {
		return nil, err
	}

	ctx := &parseContext{
		rawTokenizer: tokenizer,
	}
	qlParse(ctx)

	return ctx.parsed, ctx.err
}

func init() {
	qlErrorVerbose = true
}
