package parser

import (
	"fmt"
	"io"
)

type Tokenizer interface {
	Next() (*Token, error)

	Filename() string

	Pos() Position
}

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

func (ctx *parseContext) Lex(lval *qlSymType) int {
	token, err := ctx.Next()
	if err != nil {
		if err == io.EOF {
			return 0
		}

		ctx.err = err
		return LEX_ERROR
	}

	ctx.prevToken = token
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

func NewTokenizer(filename string, reader io.Reader) (Tokenizer, error) {
	tokenizer, err := NewRawTokenizer(filename, reader)
	if err != nil {
		return nil, err
	}

	tokenizer, err = NewCommentProcessor(tokenizer)
	if err != nil {
		return nil, err
	}

	return NewTerminatorProcessor(tokenizer)
}

func Parse(filename string, reader io.Reader) ([]Node, error) {
	tokenizer, err := NewTokenizer(filename, reader)
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
