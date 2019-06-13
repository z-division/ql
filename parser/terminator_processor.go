package parser

import (
	"io"
)

var (
	// Tokens that are sensitive to terminators on either side of the token
	terminatorSensitive = map[int]struct{}{
		IDENT:  struct{}{},
		CHAR:   struct{}{},
		STRING: struct{}{},
		INT:    struct{}{},
		FLOAT:  struct{}{},
	}

	// Tokens that are sensitive to terminators on the leading side of the
	// token
	leadingTerminatorSensitive = map[int]struct{}{
		LET: struct{}{},
	}

	// Tokens that are sensitive to terminators on the trailing side of the
	// token
	trailingTerminatorSensitive = map[int]struct{}{}

	// Tokens that are implicitly lead by a terminator
	implicitLeadingTerminator = map[int]struct{}{}

	// Tokens that are implicitly trail by a terminator
	implicitTrailingTerminator = map[int]struct{}{
		L_BRACE: struct{}{},
	}
)

// This tokenizer drops context insensitive newline tokens, and convert
// SEMICOLON and context sensitive NEWLINE into TERMINATOR
type terminatorProcessor struct {
	filename string
	base     Tokenizer

	prevToken *Token
	buffered  []*Token
}

func NewTerminatorProcessor(
	filename string,
	baseTokenizer Tokenizer) (Tokenizer, error) {

	return &terminatorProcessor{
		filename: filename,
		base:     baseTokenizer,
	}, nil
}

func (proc *terminatorProcessor) maybeFill() error {
	if len(proc.buffered) > 0 {
		return nil
	}

	var terminator *Token
	var nextNonTerminator *Token
	for {
		token, err := proc.base.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// TODO(patrick): remove hack
		if token.Type == COMMENT {
			continue
		}

		if token.Type == NEWLINE {
			if terminator == nil {
				terminator = token
			}

			continue
		}

		if token.Type == SEMICOLON {
			terminator = token
			continue
		}

		nextNonTerminator = token
		break
	}

	if proc.prevToken == nil {
		if nextNonTerminator != nil {
			proc.buffered = []*Token{nextNonTerminator}
			return nil
		}

		// "empty" file

		if terminator == nil {
			terminator = &Token{
				Location: Location{
					Filename: proc.filename,
					Start:    Position{1, 1},
					End:      Position{1, 1},
				},
			}
		}

		terminator.Type = TERMINATOR
		proc.buffered = []*Token{terminator}
		return nil
	}

	if nextNonTerminator == nil { // EOF
		if proc.prevToken.Type == TERMINATOR {
			return io.EOF
		}

		if terminator == nil {
			terminator = &Token{
				Location: proc.prevToken.Location,
			}
			terminator.Start = terminator.End
		}

		terminator.Type = TERMINATOR
		proc.buffered = []*Token{terminator}
		return nil
	}

	// Always preserve explicit terminator
	if terminator != nil && terminator.Type == SEMICOLON {
		terminator.Type = TERMINATOR
		proc.buffered = []*Token{terminator, nextNonTerminator}
		return nil
	}

	if terminator != nil {
		terminator.Type = TERMINATOR
	}

	// Insert mandatory implicit terminator
	requireTerminator := false
	_, ok := implicitTrailingTerminator[proc.prevToken.Type]
	if ok {
		requireTerminator = true
	}

	_, ok = implicitLeadingTerminator[nextNonTerminator.Type]
	if ok {
		requireTerminator = true
	}

	if requireTerminator {
		if terminator == nil {
			terminator = &Token{
				Type:     TERMINATOR,
				Location: nextNonTerminator.Location,
			}
			terminator.End = terminator.Start
		}

		proc.buffered = []*Token{terminator, nextNonTerminator}
		return nil
	}

	// No newline in between prev and next tokens.  Don't insert terminator.
	if terminator == nil {
		proc.buffered = []*Token{nextNonTerminator}
		return nil
	}

	// Preserve terminator if both prev and next token are terminator sensitive
	_, leading := leadingTerminatorSensitive[nextNonTerminator.Type]
	_, ok = terminatorSensitive[nextNonTerminator.Type]
	if ok {
		leading = true
	}

	_, trailing := trailingTerminatorSensitive[proc.prevToken.Type]
	_, ok = terminatorSensitive[proc.prevToken.Type]
	if ok {
		trailing = true
	}

	if leading && trailing {
		proc.buffered = []*Token{terminator, nextNonTerminator}
		return nil
	}

	proc.buffered = []*Token{nextNonTerminator}
	return nil
}

func (proc *terminatorProcessor) Next() (*Token, error) {
	err := proc.maybeFill()
	if err != nil {
		return nil, err
	}

	proc.prevToken = proc.buffered[0]
	proc.buffered = proc.buffered[1:]
	return proc.prevToken, nil
}
