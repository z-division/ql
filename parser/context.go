package parser

import (
	"fmt"
	"io"
)

type parseContext struct {
	filename string
	Tokenizer

	prevToken *Token

	// TODO(patrick): define parse tree
	parsed []Node

	err error
}

func (ctx *parseContext) setParsed(parsed []Node) {
	ctx.parsed = parsed
}

// TODO(patrick): implement
func (ctx *parseContext) Lex(lval *qlSymType) int {
	token, err := ctx.NextToken()
	if err != nil {
		if err == io.EOF {
			return 0
		}

		ctx.err = err
		return LEX_ERROR
	}

	ctx.prevToken = token

	// TODO(patrick): remove hack
	if token.Type == NEWLINE || token.Type == COMMENT {
		return ctx.Lex(lval)
	}

	lval.Token = token
	return token.Type
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
	tokenizer, err := NewRawTokenizer(filename, reader)
	if err != nil {
		return nil, err
	}

	ctx := &parseContext{
		filename:  filename,
		Tokenizer: tokenizer,
	}
	qlParse(ctx)

	return ctx.parsed, ctx.err
}

func init() {
	qlErrorVerbose = true
}
